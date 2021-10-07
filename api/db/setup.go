package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"simple-rest-api/api/models"
)

var db *gorm.DB

func SetupModels(dbTarget string) {
	database, err := gorm.Open("sqlite3", dbTarget)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

	db = database
}

func GetDB() *gorm.DB {
	return db
}
