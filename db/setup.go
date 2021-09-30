package db

import (
	"github.com/jinzhu/gorm"
	"simple-rest-api/models"
)

var db *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "./test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

	db = database
}

func GetDB() *gorm.DB {
	if db == nil {
		ConnectDataBase()
	}
	return db
}