package main

import (
	"context"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	authorControllers "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/controllers"
	authorRepository "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/repository"
	authorRoutes "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/routes"
	authorServices "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/services"
	bookControllers "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/controllers"
	bookRepository "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/repository"
	bookRoutes "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/routes"
	bookServices "github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/services"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/utils"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/pkg/db"
)

func main() {
	db := db.NewConnection()
	defer db.Close(context.Background())

	app := fiber.New()

	validatorUtil := utils.NewValidatorUtil(validator.New())
	querystringUtil := utils.NewQuerystringUtil()

	authorRepository := authorRepository.NewAuthorRepository(db)
	authorService := authorServices.NewAuthorService(authorRepository)
	authorController := authorControllers.NewAuthorController(authorService, querystringUtil)
	authorRouter := authorRoutes.NewAuthorRoutes(authorController)
	authorRouter.RegisterRoutes(app)

	bookRepository := bookRepository.NewBookRepository(db)
	bookService := bookServices.NewBookService(bookRepository)
	bookController := bookControllers.NewBookController(bookService, validatorUtil, querystringUtil)
	bookRouter := bookRoutes.NewBookRoutes(bookController)
	bookRouter.RegisterRoutes(app)

	log.Fatal(app.Listen(os.Getenv("BASE_URL")))
}
