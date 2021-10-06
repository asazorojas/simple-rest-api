package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	controllers2 "simple-rest-api/api/controllers"
	"simple-rest-api/api/db"
	"simple-rest-api/api/repositories"
	"simple-rest-api/api/services"
)

func getGinEngine() *gin.Engine {
	return gin.Default()
}

func Start() {
	db.ConnectDataBase()

	router := getGinEngine()

	setupV100Routes(router)
	setupV101Routes(router)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8080")
}

func setupV100Routes(router *gin.Engine) {
	apiV100 := router.Group("/api/v100")
	{
		apiV100.GET("/books", controllers2.FindBooks)
		apiV100.POST("/books", controllers2.CreateBook)
		apiV100.GET("/books/:id", controllers2.FindBook)
		apiV100.PATCH("/books/:id", controllers2.UpdateBook)
		apiV100.DELETE("/books/:id", controllers2.DeleteBook)
	}
}

func setupV101Routes(router *gin.Engine) {
	booksRepository := repositories.BookRepository{DB: db.GetDB()}
	booksService := services.BooksService{IBookRepository: booksRepository}
	booksController := controllers2.BooksController{IBooksService : booksService}

	apiV101 := router.Group("/api/v101")
	{
		apiV101.GET("/books", booksController.FindBooksV2)
	}
}
