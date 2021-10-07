Please use Golang 1.16.4 from the beginning because there seems to be a problem with Go1.17 and gin gonic.

If you already have a golang 1.17.x please do as follows:

```
go get golang.org/dl/go1.16.4
go1.16.4 download
```

Execute the api running:

Newest folder structure:

```
go1.16.4 run main.go
```

Old folder structure:

```
Go1.16.4 run api/api.go
```

To remove the vendor folder from the project just delete the vendor folder with ``rm -rf vendor`` and then run:

```
go1.16.4 mod tidy
```

Extra dependencies if needed are added with go get url of the dependencies and then re run ``go1.16.4 mod tidy`` again.

To re-create the vendor folder in your project just run ``go1.16.4 mod vendor``

To run tests, when available just run:

```
go1.16.4 test ./...
```

## Dependencies

> GoSqlMock : github.com/DATA-DOG/go-sqlmock v1.5.0

> Gin: github.com/gin-gonic/gin v1.7.4

> GORM: github.com/jinzhu/gorm v1.9.16

> Testify: github.com/stretchr/testify v1.4.0 
