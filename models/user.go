package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID      uint   `json:"id"	gorm:"primary_key"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
