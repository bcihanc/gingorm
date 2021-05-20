package main

import (
	"gingorm/controllers"
	"gingorm/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{"admin": "password"}))

	authorized.GET("/user", controllers.GetUsers)

	r.GET("/user/:id", controllers.GetUserById)
	r.POST("/user", controllers.CreateUser)
	r.POST("/user/update/:id", controllers.UpdateUser)
	r.GET("/user/delete/:id", controllers.DeleteUser)
	r.Run()
}
