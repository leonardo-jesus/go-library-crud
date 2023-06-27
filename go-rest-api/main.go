package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/controllers"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/repository"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/routes"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/services"
)

const BASE_URL string = "localhost:3333"

var authors = []*models.Author{
	{
		Id:   1,
		Name: "Luciano Ramalho",
	},
	{
		Id:   2,
		Name: "Osvaldo Santana Neto",
	},
	{
		Id:   3,
		Name: "David Beazley",
	},
	{
		Id:   4,
		Name: "Chetan Giridhar",
	},
}

func main() {
	app := fiber.New()

	authorRepository := repository.NewAuthorRepository(authors)
	authorService := services.NewAuthorService(authorRepository)
	authorController := controllers.NewAuthorController(authorService)
	authorRouter := routes.NewAuthorRoutes(authorController)
	authorRouter.RegisterRoutes(app)

	log.Fatal(app.Listen(BASE_URL))
}
