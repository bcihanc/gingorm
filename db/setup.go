package db

import (
	"fmt"
	"gingorm/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		fmt.Printf("Error connecting to DB: %s", err)
	}
	database.AutoMigrate(&models.User{})
	DB = database
}
