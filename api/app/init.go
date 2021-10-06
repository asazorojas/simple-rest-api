package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/api/controllers"
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
		apiV100.GET("/books", controllers.FindBooks)
		apiV100.POST("/books", controllers.CreateBook)
		apiV100.GET("/books/:id", controllers.FindBook)
		apiV100.PATCH("/books/:id", controllers.UpdateBook)
		apiV100.DELETE("/books/:id", controllers.DeleteBook)
	}
}

func setupV101Routes(router *gin.Engine) {
	booksRepository := repositories.BookRepository{DB: db.GetDB()}
	booksService := services.BooksService{IBookRepository: booksRepository}
	booksController := controllers.BooksV101Controller{IBooksService : booksService}

	apiV101 := router.Group("/api/v101")
	{
		apiV101.GET("/books", booksController.FindBooksV2)
	}
}
