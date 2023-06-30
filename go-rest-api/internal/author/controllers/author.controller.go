package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/services"
)

type AuthorControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByName(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type authorController struct {
	authorService services.AuthorServiceInterface
}

func NewAuthorController(authorService services.AuthorServiceInterface) AuthorControllerInterface {
	return &authorController{authorService}
}

func (c *authorController) FindAll(ctx *fiber.Ctx) error {
	page := c.GetPageFromQuerystring(ctx)

	authors, err := c.authorService.FindAll(page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(authors)
}

func (c *authorController) FindByName(ctx *fiber.Ctx) error {
	page := c.GetPageFromQuerystring(ctx)

	authors, err := c.authorService.FindByName(ctx.Query("name"), page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(authors)
}

func (c *authorController) Create(ctx *fiber.Ctx) error {
	err := c.authorService.Create()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"success": "true"})
}

func (c *authorController) GetPageFromQuerystring(ctx *fiber.Ctx) int {
	result, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return 1
	}

	if result < 1 {
		result = 1
	}

	return result
}
