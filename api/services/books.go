package services

import (
	"simple-rest-api/api/dtos"
	"simple-rest-api/api/interfaces"
)

type BooksService struct {
	interfaces.IBookRepository
}

func (booksService BooksService) GetBooks() ([]dtos.BookDataV2, error) {
	books, err := booksService.IBookRepository.GetBooks()
	var booksData []dtos.BookDataV2

	if err != nil {
		return nil, err
	}

	if len(books) > 0 {
		for _, book := range books {
			booksData = append(booksData, dtos.BookDataV2{
				ID:     book.ID,
				Author: book.Author,
				Title:  book.Title,
			})
		}
	} else {
		booksData = []dtos.BookDataV2{}
	}
	return booksData, err
}

func (booksService BooksService) GetBookById(bookId int) (dtos.BookDataV2, error) {
	book, err := booksService.IBookRepository.GetBookById(bookId)
	return dtos.BookDataV2{
		ID:     book.ID,
		Author: book.Author,
		Title:  book.Title,
	}, err
}
