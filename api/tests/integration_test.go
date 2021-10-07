package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simple-rest-api/api/app"
	"simple-rest-api/api/db"
	"simple-rest-api/api/models"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestBooksCrud(t *testing.T) {
	dbTarget := "test.db"

	db.SetupModels(dbTarget)
	dbInstance := db.GetDB()
	dbInstance.DropTableIfExists(&models.Book{}, "books")
	db.SetupModels(dbTarget)
	dbInstance = db.GetDB()
	defer dbInstance.Close()

	router := gin.Default()
	app.SetupV100RoutesMapping(router)
	app.SetupV101RoutesMapping(router)

	gin.SetMode(gin.TestMode)

	t.Run("Create Empty DB", func(t *testing.T) {
		w := performRequest(router, "GET", "/api/v100/books")

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Retrieve Nonexistent ID on Empty DB", func(t *testing.T) {

		w := performRequest(router, "GET", "/api/v100/books/2")

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Populate DB with Harry Potter Set", func(t *testing.T) {
		books := []string{
			"Harry Potter and The Philosopher's Stone",
			"Harry Potter and The Chamber of Secrets",
			"Harry Potter and The Prisoner of Azkaban",
			"Harry Potter and The Goblet of Fire",
			"Harry Potter and The Order of The Phoenix",
			"Harry Potter and The Half-Blood Prince",
			"Harry Potter and The Deathly Hallows",
		}

		for _, book := range books {

			payload, _ := json.Marshal(map[string]interface{}{
				"author": "J. K. Rowling",
				"title":  book,
			})

			req, err := http.NewRequest("POST", "/api/v100/books", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, nil, err)
			assert.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Retrieve Existing ID on Populated DB", func(t *testing.T) {
		w := performRequest(router, "GET", "/api/v100/books/2")

		expected := V1Response{
			Data: V1ResponseData{
				ID		:  2,
				Title	: "Harry Potter and The Chamber of Secrets",
				Author	: "J. K. Rowling",
			},
		}

		var response V1Response
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, response)
	})

	t.Run("Retrieve Existing ID on Populated DB V101", func(t *testing.T) {
		w := performRequest(router, "GET", "/api/v101/books/2")

		expected := V2Response{
			Data: V2ResponseData{
				ID		:  2,
				Title	: "Harry Potter and The Chamber of Secrets",
				Author	: "J. K. Rowling",
			},
		}

		var response V2Response
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, response)
	})
}

type V1ResponseData struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}
type V1Response struct {
	Data V1ResponseData `json:"data"`
}

type V2ResponseData struct {
	ID int `json:"book_id"`
	Title string `json:"book_title"`
	Author string `json:"book_author"`
}
type V2Response struct {
	Data V2ResponseData `json:"data"`
}
