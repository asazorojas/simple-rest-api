package repositories

import (
	"github.com/jinzhu/gorm"
	"simple-rest-api/models"
)

type BookRepository struct {
	DB *gorm.DB
}

func (bookRepository BookRepository) GetBooks() ([]models.Book, error) {
	return nil, nil
}

func (bookRepository BookRepository) GetBookById(bookId int) (models.Book, error) {
	return models.Book{}, nil
}
