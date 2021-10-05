package services

import (
	"simple-rest-api/dtos"
	"simple-rest-api/interfaces"
)

type BooksService struct {
	interfaces.IBookRepository
}

func (booksService BooksService) GetBooks() ([]dtos.BookData, error) {
	books, err := booksService.IBookRepository.GetBooks()
	var booksData []dtos.BookData
	for _, book := range books {
		booksData = append(booksData, dtos.BookData{
			ID:     book.ID,
			Author: book.Author,
			Title:  book.Title,
		})
	}
	return booksData, err
}

func (booksService BooksService) GetBookById(bookId int) (dtos.BookData, error) {
	book, err := booksService.IBookRepository.GetBookById(bookId)
	return dtos.BookData{
		ID:     book.ID,
		Author: book.Author,
		Title:  book.Title,
	}, err
}
