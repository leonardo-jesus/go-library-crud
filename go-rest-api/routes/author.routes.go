package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/controllers"
)

type authorRoutes struct {
	authorController controllers.AuthorControllerInterface
}

func NewAuthorRoutes(authorController controllers.AuthorControllerInterface) RoutesInterface {
	return &authorRoutes{authorController}
}

func (r *authorRoutes) RegisterRoutes(app *fiber.App) {
	app.Get("/", r.authorController.Check)
	app.Get("/author", r.authorController.FindAll)
}
