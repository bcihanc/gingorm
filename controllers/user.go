package controllers

import (
	"gingorm/db"
	"gingorm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserInput struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input createUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Surname: input.Surname}
	db.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

