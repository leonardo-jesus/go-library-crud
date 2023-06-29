package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/controllers"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/repository"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/routes"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/services"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/pkg/db"
)

func main() {
	db := db.NewConnection()
	defer db.Close(context.Background())

	app := fiber.New()

	authorRepository := repository.NewAuthorRepository(db)
	authorService := services.NewAuthorService(authorRepository)
	authorController := controllers.NewAuthorController(authorService)
	authorRouter := routes.NewAuthorRoutes(authorController)
	authorRouter.RegisterRoutes(app)

	log.Fatal(app.Listen(os.Getenv("BASE_URL")))
}
