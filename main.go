package main

import (
	"gingorm/controllers"
	"gingorm/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World."})
	})
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.Run()
}
