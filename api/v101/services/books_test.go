package servicesV101

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"simple-rest-api/api/dtos"
	"simple-rest-api/api/models"
	"testing"

	"github.com/stretchr/testify/mock"
)

type bookRepositoryMock struct {
	mock.Mock
}

func (r bookRepositoryMock) GetBooks() ([]*models.Book, error) {
	args := r.Called()
	return args.Get(0).([]*models.Book), args.Error(1)
}

func (r bookRepositoryMock) GetBookById(bookId int) (*models.Book, error) {
	args := r.Called(bookId)
	return args.Get(0).(*models.Book), args.Error(1)
}

func getMockedBooks() []*models.Book {
	var mockedBooks []*models.Book
	mockedBooks = append(mockedBooks, &models.Book{
		ID: 1, Title: "The Lord of the Rings", Author: "JRR Tolkien",
	})
	mockedBooks = append(mockedBooks, &models.Book{
		ID: 1, Title: "The Hobbit", Author: "JRR Tolkien",
	})
	return mockedBooks
}

func TestGetAllBooks(t *testing.T) {
	bookRepositoryMock := new(bookRepositoryMock)
	bookRepositoryMock.On("GetBooks").Return(getMockedBooks(), nil)

	booksService := BooksService{IBookRepository: bookRepositoryMock}

	books, err := booksService.GetBooks()

	bookRepositoryMock.AssertExpectations(t)
	assert.NotNil(t, books)
	assert.Nil(t, err)
}

func TestGetBookByIdWhenBookExists(t *testing.T) {
	bookId := 1
	expectedBook := &dtos.BookDataV2{
		ID:     1,
		Title:  "The Lord of the Rings",
		Author: "JRR Tolkien",
	}
	bookRepositoryMock := new(bookRepositoryMock)

	bookRepositoryMock.On("GetBookById", bookId).Return(&models.Book{
		ID:     1,
		Title:  "The Lord of the Rings",
		Author: "JRR Tolkien",
	}, nil)

	booksService := BooksService{IBookRepository: bookRepositoryMock}

	book, err := booksService.GetBookById(bookId)

	bookRepositoryMock.AssertExpectations(t)
	assert.NotNil(t, book)
	assert.Nil(t, err)
	assert.Equal(t, expectedBook, book)
}

func TestGetBookByIdWhenBookDoesNotExist(t *testing.T) {
	bookId := 1
	bookRepositoryMock := new(bookRepositoryMock)

	bookRepositoryMock.On("GetBookById", bookId).Return(
		(*models.Book)(nil), errors.New("Record not found"))

	booksService := BooksService{IBookRepository: bookRepositoryMock}

	book, err := booksService.GetBookById(bookId)

	bookRepositoryMock.AssertExpectations(t)
	assert.Nil(t, book)
	assert.NotNil(t, err)
	assert.Equal(t, "Record not found", err.Error())
}
