package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.Run()
}


























/* //With custom things
func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
*/


/* //Custom log format
func main() {
	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}*/


/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/controllers"
	"simple-rest-api/models"
)

func main() {
	router := gin.Default()
	//apiV100 := router.Group("/api/v100")
	//{
	//	apiV100.GET("/books", controllers.FindBooks)
	//	apiV100.POST("/books", controllers.CreateBook)
	//	apiV100.GET("/books/:id", controllers.FindBook)
    //	apiV100.PATCH("/books/:id", controllers.UpdateBook)
	//	apiV100.DELETE("/books/:id", controllers.DeleteBook)
	//}

	db.ConnectDataBase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
    router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8080")
}
*/

//booksRepository := repositories.BookRepository{DB: db.GetDB()}
//booksService := services.BooksService{IBookRepository: booksRepository}
//booksController := controllers.BooksController{IBooksService : booksService}

//router.GET("/ping", booksController.FindBooks)