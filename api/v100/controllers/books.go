package controllersV100

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/api/db"
	"simple-rest-api/api/models"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	db.GetDB().Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	db.GetDB().Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book

	bookId := c.Param("id")
	if err := db.GetDB().Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	bookId := c.Param("id")
	if err := db.GetDB().Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.GetDB().Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	bookId := c.Param("id")
	if err := db.GetDB().Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.GetDB().Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}