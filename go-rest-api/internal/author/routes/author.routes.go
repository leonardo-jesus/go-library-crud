package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/controllers"
)

type authorRoutes struct {
	authorController controllers.AuthorControllerInterface
}

func NewAuthorRoutes(authorController controllers.AuthorControllerInterface) *authorRoutes {
	return &authorRoutes{authorController}
}

func (r *authorRoutes) RegisterRoutes(app *fiber.App) {
	app.Get("/", r.authorController.Check)
	app.Get("/author", r.authorController.FindAll)
	app.Get("/author/filter", r.authorController.FindByName)
	app.Post("/author", r.authorController.Create)
}
