package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/services"
)

type AuthorControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	Check(c *fiber.Ctx) error
}

type authorController struct {
	authorService services.AuthorServiceInterface
}

func NewAuthorController(authorService services.AuthorServiceInterface) AuthorControllerInterface {
	return &authorController{authorService}
}

func (c *authorController) FindAll(ctx *fiber.Ctx) error {
	authors, err := c.authorService.FindAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(authors)
}

func (c *authorController) Check(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"success": "true"})
}
