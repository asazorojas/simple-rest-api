package controllersV101

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"simple-rest-api/api/dtos"
	"strconv"
	"testing"
)

func AssertExpected(t *testing.T, expected, received interface{}) {
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("Got %v, wanted %v", received, expected)
	}
}

func SetupContext() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return w, c
}

func SetupRequestBody(c *gin.Context, payload interface{}) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(payload)

	c.Request = &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes.Bytes())),
	}
}

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

func getMockedBooks() []*dtos.BookDataV2 {
	var mockedBooks []*dtos.BookDataV2
	mockedBooks = append(mockedBooks, &dtos.BookDataV2{
		ID: 1, Title: "The Lord of the Rings", Author: "JRR Tolkien",
	})
	mockedBooks = append(mockedBooks, &dtos.BookDataV2{
		ID: 2, Title: "The Hobbit", Author: "JRR Tolkien",
	})
	return mockedBooks
}

func TestCRUDFunctions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Get all books from DB", func(t *testing.T) {
		expectedResponse := AllBooksResponse{
			Books: []*dtos.BookDataV2{
				{
					ID: 1, Title: "The Lord of the Rings", Author: "JRR Tolkien",
				},
				{
					ID: 2, Title: "The Hobbit", Author: "JRR Tolkien",
				},
			},
		}
		booksServiceMock := new(bookServiceMock)
		booksServiceMock.On("GetBooks").Return(getMockedBooks(), nil)
		w, c := SetupContext()

		b := BooksV101Controller{IBooksService: booksServiceMock}
		b.FindBooksV2(c)

		var response AllBooksResponse
		json.Unmarshal([]byte(w.Body.String()), &response)

		booksServiceMock.AssertExpectations(t)
		AssertExpected(t, w.Code, http.StatusOK)
		AssertExpected(t, len(w.Body.String()) > 0, true)
		AssertExpected(t, expectedResponse, response)
	})

	t.Run("Get one book by id from DB and the book exists", func(t *testing.T) {
		id := 1
		expectedResponse := BookByIdV2Response{
			Data: dtos.BookDataV2{ID: 1, Title: "The Lord of the Rings", Author: "JRR Tolkien"},
		}
		booksServiceMock := new(bookServiceMock)
		booksServiceMock.On("GetBookById", id).Return(&dtos.BookDataV2{
			ID: 1, Title: "The Lord of the Rings", Author: "JRR Tolkien",
		}, nil)
		w, c := SetupContext()
		c.Params = []gin.Param{gin.Param{Key: "id", Value: strconv.Itoa(id)}}

		b := BooksV101Controller{IBooksService: booksServiceMock}
		b.FindBookByIdV2(c)

		var response BookByIdV2Response
		json.Unmarshal([]byte(w.Body.String()), &response)

		booksServiceMock.AssertExpectations(t)
		AssertExpected(t, w.Code, http.StatusOK)
		AssertExpected(t, len(w.Body.String()) > 0, true)
		AssertExpected(t, expectedResponse, response)
	})

	t.Run("Get one book by id from DB and the book does not exist", func(t *testing.T) {
		id := 100
		expectedResponse := ErrorResponse{
			Error: "Record not found",
		}
		booksServiceMock := new(bookServiceMock)
		booksServiceMock.On("GetBookById", id).Return((*dtos.BookDataV2)(nil), errors.New("Record not found"))
		w, c := SetupContext()
		c.Params = []gin.Param{gin.Param{Key: "id", Value: strconv.Itoa(id)}}

		b := BooksV101Controller{IBooksService: booksServiceMock}
		b.FindBookByIdV2(c)

		var response ErrorResponse
		json.Unmarshal([]byte(w.Body.String()), &response)

		booksServiceMock.AssertExpectations(t)
		AssertExpected(t, w.Code, http.StatusBadRequest)
		AssertExpected(t, len(w.Body.String()) > 0, true)
		AssertExpected(t, expectedResponse, response)
	})
}

type BookByIdV2Response struct {
	Data dtos.BookDataV2 `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
