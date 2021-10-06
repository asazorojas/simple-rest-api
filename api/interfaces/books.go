package interfaces

import (
	"simple-rest-api/api/dtos"
	"simple-rest-api/api/models"
)

type IBookRepository interface {
	GetBooks() ([]models.Book, error)
	GetBookById(bookId int) (models.Book, error)
}

type IBooksService interface {
	GetBooks() ([]dtos.BookDataV2, error)
	GetBookById(bookId int) (dtos.BookDataV2, error)
}