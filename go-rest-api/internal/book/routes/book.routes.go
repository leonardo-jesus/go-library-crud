package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/controllers"
)

type BookRoutesInterface interface {
	RegisterRoutes(app *fiber.App)
}

type bookRoutes struct {
	bookController controllers.BookControllerInterface
}

func NewBookRoutes(bookController controllers.BookControllerInterface) *bookRoutes {
	return &bookRoutes{bookController}
}

func (r *bookRoutes) RegisterRoutes(app *fiber.App) {
	app.Get("/book", r.bookController.FindAll)
	app.Get("/book/filter", r.bookController.FindByFilteredBooks)
	app.Post("/book", r.bookController.Create)
	app.Patch("/book/:id", r.bookController.Update)
	app.Delete("/book/:id", r.bookController.Delete)
}
