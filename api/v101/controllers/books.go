package controllersV101

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/api/dtos"
	"simple-rest-api/api/interfaces"
	"strconv"
)

type BooksV101Controller struct {
	interfaces.IBooksService
}

func (controller BooksV101Controller) FindBooksV2(c *gin.Context) {
	books, err := controller.IBooksService.GetBooks()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if books != nil && len(books) > 0 {
		c.JSON(http.StatusOK, AllBooksResponse{
			Books: books,
		})
		return
	}

	c.JSON(http.StatusOK, AllBooksResponse{
		Books: []*dtos.BookDataV2{},
	})

}

type AllBooksResponse struct {
	Books []*dtos.BookDataV2 `json:"books"`
}

func (controller BooksV101Controller) FindBookByIdV2(c *gin.Context) {
	id, idErr := strconv.ParseInt(c.Param("id"), 10, 0)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": idErr.Error()})
		return
	}
	book, err := controller.IBooksService.GetBookById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})

}
