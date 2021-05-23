package db

import (
	"errors"
	"gingorm/models"
	"gingorm/utils"
)

func LoginWithEmailAndPassword(user *models.User) (bool, error) {
	var userFromDB models.User
	findUserDB := DB.Where("email = ?", user.Email).First(&userFromDB)
	if findUserDB.Error != nil {
		return false, findUserDB.Error
	}

	isPasswordTrue := utils.CheckPasswordHash(user.Password, userFromDB.Password)

	if isPasswordTrue {
		return true, nil
	} else {
		return false, errors.New("password is wrong")
	}
}
