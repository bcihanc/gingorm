package controllers

import (
	"gingorm/db"
	"gingorm/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthenticatorHandler(c *gin.Context) (interface{}, error) {
	var loginInput models.UserLogin
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	user := models.User{Email: loginInput.Email, Password: loginInput.Password}

	isPasswordTrue, loginErr := db.LoginWithEmailAndPassword(&user)

	if loginErr != nil {
		log.Print(loginErr.Error())
		return nil, loginErr
	}

	if isPasswordTrue {
		return &user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
