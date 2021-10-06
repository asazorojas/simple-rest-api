package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/api/db"
	controllersV100 "simple-rest-api/api/v100/controllers"
	controllersV101 "simple-rest-api/api/v101/controllers"
	repositoriesV101 "simple-rest-api/api/v101/repositories"
	servicesV101 "simple-rest-api/api/v101/services"
)

func getGinEngine() *gin.Engine {
	return gin.New()
}

func Start() {
	// Try to connect to DB
	db.ConnectDataBase()

	// Get Gin Engine
	router := getGinEngine()

	// Set logger and recovery middleware
	addLoggerMiddleware(router)
	addRecoveryMiddleware(router)

	// Setup routes for V100 and V101 endpoints
	setupV100RoutesMapping(router)
	setupV101RoutesMapping(router)
	setupNoRouteMapping(router)

	router.Run(":8080")
}

func setupNoRouteMapping(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func addLoggerMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
}

func addRecoveryMiddleware(router *gin.Engine) {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
}

func setupV100RoutesMapping(router *gin.Engine) {
	apiV100 := router.Group("/api/v100")
	{
		apiV100.GET("/books", controllersV100.FindBooks)
		apiV100.POST("/books", controllersV100.CreateBook)
		apiV100.GET("/books/:id", controllersV100.FindBook)
		apiV100.PATCH("/books/:id", controllersV100.UpdateBook)
		apiV100.DELETE("/books/:id", controllersV100.DeleteBook)
	}
}

func setupV101RoutesMapping(router *gin.Engine) {
	booksRepository := repositoriesV101.BookRepository{DB: db.GetDB()}
	booksService := servicesV101.BooksService{IBookRepository: booksRepository}
	booksController := controllersV101.BooksV101Controller{IBooksService: booksService}

	apiV101 := router.Group("/api/v101")
	{
		apiV101.GET("/books", booksController.FindBooksV2)
	}
}
