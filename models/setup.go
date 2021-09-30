package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "./test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	db = database
}

func GetDB() *gorm.DB {
	return db
}
