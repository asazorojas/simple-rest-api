package repositoriesV101

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

//type Suite struct {
//	suite.Suite
//	DB   *gorm.DB
//	mock sqlmock.Sqlmock
//
//	repository BookRepository
//	book     *models.Book
//}
//
//func (s *Suite) SetupSuite() {
//	var (
//		db  *sql.DB
//		err error
//	)
//
//	db, s.mock, err = sqlmock.New()
//	require.NoError(s.T(), err)
//
//	s.DB, err = gorm.Open("sqlite3", db)
//	require.NoError(s.T(), err)
//
//	s.DB.LogMode(true)
//
//	s.repository = BookRepository{DB: s.DB}
//}
//
//func (s *Suite) AfterTest(_, _ string) {
//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
//}
//
//func TestInit(t *testing.T) {
//	suite.Run(t, new(Suite))
//}
//
//func (s *Suite) Test_repository_GetAllBooks() {
//	var (
//		id = 1
//		title = "The Lord of the Rings"
//		author = "J.R.R Tolkien"
//	)
//
//	s.mock.ExpectQuery(
//		regexp.QuoteMeta(`SELECT * FROM "books"`)).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "author", "title"}).
//			AddRow(id, author, title))
//
//	books, err := s.repository.GetBooks()
//
//	require.NoError(s.T(), err)
//	assert.NotNil(s.T(), books)
//}

func Test_repository_GetAllBooks(t *testing.T) {
	var (
		id = 1
		title = "The Lord of the Rings"
		author = "J.R.R Tolkien"
	)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("sqlite3", db)
	gdb.LogMode(true)
	repository := BookRepository{DB: gdb}

	mock.ExpectQuery(
		"SELECT * FROM \"books\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "author", "title"}).
			AddRow(id, author, title))

	books, err := repository.GetBooks()

	require.NoError(t, err)
	assert.NotNil(t, books)
}


func Test_repository_GetOneBook(t *testing.T) {
	var (
		id = 1
		title = "The Lord of the Rings"
		author = "J.R.R Tolkien"
	)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("sqlite3", db)
	repository := BookRepository{DB: gdb}

	sql := `SELECT * FROM "books" WHERE (id = ?) ORDER BY "books"."id" ASC LIMIT 1`

	mock.ExpectQuery(sql).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "author", "title"}).
			AddRow(id, author, title))

	book, err := repository.GetBookById(id)

	require.NoError(t, err)
	assert.NotNil(t, book)
}
