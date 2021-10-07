package repositoriesV101

import (
	"github.com/jinzhu/gorm"
	"simple-rest-api/api/models"
)

type BookRepository struct {
	DB *gorm.DB
}

func (bookRepository BookRepository) GetBooks() ([]*models.Book, error) {
	var books []*models.Book
	err := bookRepository.DB.Find(&books).Error
	return books, err
}

func (bookRepository BookRepository) GetBookById(bookId int) (*models.Book, error) {
	var book models.Book
	err := bookRepository.DB.First(&book, "id = ?", bookId).Error
	// All these forms are equivalent to SELECT * FROM books WHERE books.id = bookId ORDER BY books.id ASC LIMIT 1
	//err := bookRepository.DB.Where("id = ?", bookId).First(&book).Error
	//err := bookRepository.DB.First(&book, bookId).Error
	return &book, err
}
