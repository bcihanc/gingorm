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
	db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db.DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	var input createUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.First(&user, id)

	user.Name = input.Name
	user.Surname = input.Surname

	db.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result := db.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "user deleted from db"})
}
