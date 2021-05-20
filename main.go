package main

import (
	"gingorm/controllers"
	"gingorm/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{"admin": "password"}))

	authorized.GET("/user", controllers.GetUsers)

	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUserById)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	err := r.Run()

	if err != nil {
		log.Panicln("gin failed", err)
	}
}
