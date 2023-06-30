package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/controllers"
)

type bookRoutes struct {
	bookController controllers.BookControllerInterface
}

func NewBookRoutes(bookController controllers.BookControllerInterface) *bookRoutes {
	return &bookRoutes{bookController}
}

func (r *bookRoutes) RegisterRoutes(app *fiber.App) {
	app.Get("/book", r.bookController.FindAll)
	app.Get("/book/filter", r.bookController.FindByName)
	app.Post("/book", r.bookController.Create)
}
