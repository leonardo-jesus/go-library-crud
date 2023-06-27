package routes

import "github.com/gofiber/fiber/v2"

type RoutesInterface interface {
	RegisterRoutes(app *fiber.App)
}
