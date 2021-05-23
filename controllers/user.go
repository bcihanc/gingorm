package controllers

import (
	"gingorm/db"
	"gingorm/models"
	"gingorm/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func AuthorizedData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "authorized"})
}

func CreateUser(c *gin.Context) {
	var input models.UserLogin
	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}
	hashedPassword, _ := utils.HashPassword(input.Password)
	user := models.User{Email: input.Email, Password: hashedPassword}
	db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db.DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginWithEmailAndPasswordEndPoint(c *gin.Context) {
	var loginInput models.UserLogin
	if bindErr := c.ShouldBindJSON(&loginInput); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	user := models.User{Email: loginInput.Email, Password: loginInput.Password}
	isPasswordTrue, loginErr := db.LoginWithEmailAndPassword(&user)

	if loginErr != nil {
		c.JSON(http.StatusOK, gin.H{"error": loginErr.Error()})
		return
	}

	if isPasswordTrue {
		c.JSON(http.StatusOK, gin.H{"data": loginInput})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"data": "password is wrong"})
	}

}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	var input models.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.First(&user, id)

	user.Email = input.Email
	user.Password = input.Password

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
