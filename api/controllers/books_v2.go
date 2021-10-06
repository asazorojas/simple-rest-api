package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/api/interfaces"
)

type BooksController struct {
	interfaces.IBooksService
}

func (controller BooksController) FindBooksV2(c *gin.Context) {
	books, err := controller.IBooksService.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
