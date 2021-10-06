package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/api/dtos"
	"simple-rest-api/api/interfaces"
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
	c.JSON(http.StatusOK, AllBooksResponse{
		Books: books,
	})
}

type AllBooksResponse struct {
	Books []dtos.BookDataV2 `json:"books"`
}
