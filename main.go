package main

import (
	"gingorm/controllers"
	"gingorm/db"
	"gingorm/docs"
	"gingorm/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	// swagger settings
	docs.SwaggerInfo.Title = "Swagger Example API writter by Go"
	docs.SwaggerInfo.Description = "This is a sample server session server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// session & swagger
		sessionGroup := v1.Group("/session")
		{
			store := cookie.NewStore([]byte("secret"))
			sessionGroup.Use(sessions.Sessions("my_session", store))
			sessionGroup.POST("", controllers.CreateVariableInSession)
			sessionGroup.GET(":key", controllers.GetVariableFromSession)
		}

		// postgres
		userGroup := v1.Group("/user")
		{
			db.ConnectDatabase()
			userGroup.POST("register", controllers.CreateUser)
			userGroup.GET("login", controllers.LoginWithEmailAndPasswordEndPoint)
			userGroup.GET(":id", controllers.GetUserById)
			userGroup.PUT(":id", controllers.UpdateUser)
			userGroup.DELETE(":id", controllers.DeleteUser)
		}

		// basic auth
		authorized := v1.Group("/auth", gin.BasicAuth(gin.Accounts{"admin": "password"}))
		{
			authorized.GET("", controllers.AuthorizedData)
		}

		// jwt auth
		identityKey := "id"

		authMiddleware, jwtErr := jwt.New(&jwt.GinJWTMiddleware{
			Realm:       "test zone",
			Key:         []byte("secret key"),
			Timeout:     time.Hour,
			MaxRefresh:  time.Hour,
			IdentityKey: identityKey,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*models.User); ok {
					return jwt.MapClaims{
						identityKey: v.Email,
					}
				}
				return jwt.MapClaims{}
			},
			IdentityHandler: func(c *gin.Context) interface{} {
				claims := jwt.ExtractClaims(c)
				return &models.User{
					Email: claims[identityKey].(string),
				}
			},
			Authenticator: controllers.AuthenticatorHandler,
			//Authorizator: func(data interface{}, c *gin.Context) bool {
			//	if v, ok := data.(*models.User); ok && v.Email == "admin" {
			//		return true
			//	}
			//
			//	return false
			//},
			Unauthorized: func(c *gin.Context, code int, message string) {
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			},
			// TokenLookup is a string in the form of "<source>:<name>" that is used
			// to extract token from the request.
			// Optional. Default value "header:Authorization".
			// Possible values:
			// - "header:<name>"
			// - "query:<name>"
			// - "cookie:<name>"
			// - "param:<name>"
			TokenLookup: "header: Authorization, query: token, cookie: jwt",
			// TokenLookup: "query:token",
			// TokenLookup: "cookie:token",

			// TokenHeadName is a string in the header. Default value is "Bearer"
			TokenHeadName: "Bearer",

			// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
			TimeFunc: time.Now,
		})

		if jwtErr != nil {
			log.Fatal("JWT Error:" + jwtErr.Error())
		}

		if errInit := authMiddleware.MiddlewareInit(); errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}

		auth := r.Group("/auth")
		auth.POST("jwt", authMiddleware.LoginHandler)

		// Refresh time can be longer than token timeout
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		auth.Use(authMiddleware.MiddlewareFunc())
	}

	if err := r.Run(); err != nil {
		log.Panicln("gin failed", err)
	}
}
