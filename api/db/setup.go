package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"simple-rest-api/api/models"
)

var db *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

	db = database
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Println("it was nil")
		ConnectDataBase()
	}
	log.Println("it is nil no more, returning...")
	return db
}
