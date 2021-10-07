package controllersV101

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-rest-api/api/dtos"
	"testing"
)

type bookServiceMock struct {
	mock.Mock
}

func (s bookServiceMock) GetBooks() ([]*dtos.BookDataV2, error) {
	args := s.Called()
	return args.Get(0).([]*dtos.BookDataV2), args.Error(1)
}

func (s bookServiceMock) GetBookById(bookId int) (*dtos.BookDataV2, error) {
	args := s.Called(bookId)
	return args.Get(0).(*dtos.BookDataV2), args.Error(1)
}

func TestGetAllBooksWhenNoRecordsInDBSoResponseIsEmpty(t *testing.T) {
	bookServiceMock := new(bookServiceMock)
	bookServiceMock.On("GetBooks").Return([]*dtos.BookDataV2{}, nil)

	booksController := BooksV101Controller{IBooksService: bookServiceMock}

	books, err := booksController.GetBooks()

	assert.NotNil(t, books)
	assert.Empty(t, books)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(books))
}

func TestGetAllBooksWhenNoRecordsInDBSoResponseIsNotEmpty(t *testing.T) {
	bookServiceMock := new(bookServiceMock)
	bookServiceMock.On("GetBooks").Return(getMockedBooks(), nil)

	booksController := BooksV101Controller{IBooksService: bookServiceMock}

	books, err := booksController.GetBooks()

	assert.NotNil(t, books)
	assert.NotEmpty(t, books)
	assert.Nil(t, err)
	assert.Greater(t, len(books), 0)
}

func TestGetAllBooksWhenErrorSoResponseIsEmpty(t *testing.T) {
	bookServiceMock := new(bookServiceMock)
	bookServiceMock.On("GetBooks").Return(([]*dtos.BookDataV2)(nil), errors.New("There was an error retrieving the books"))

	booksController := BooksV101Controller{IBooksService: bookServiceMock}

	books, err := booksController.GetBooks()

	assert.Nil(t, books)
	assert.Empty(t, books)
	assert.NotNil(t, err)
	assert.Equal(t, "There was an error retrieving the books", err.Error())
}

func getMockedBooks() []*dtos.BookDataV2 {
	var mockedBooks []*dtos.BookDataV2
	mockedBooks = append(mockedBooks, &dtos.BookDataV2{
		ID: 1, Title: "The Lord of the Rings", Author: "JRR Tolkien",
	})
	mockedBooks = append(mockedBooks, &dtos.BookDataV2{
		ID: 1, Title: "The Hobbit", Author: "JRR Tolkien",
	})
	return mockedBooks
}