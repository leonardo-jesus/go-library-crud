package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/services"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/utils"
)

type AuthorControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByName(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type authorController struct {
	authorService   services.AuthorServiceInterface
	querystringUtil utils.QuerystringUtilInterface
}

func NewAuthorController(authorService services.AuthorServiceInterface, querystringUtil utils.QuerystringUtilInterface) AuthorControllerInterface {
	return &authorController{authorService, querystringUtil}
}

func (c *authorController) FindAll(ctx *fiber.Ctx) error {
	page := c.querystringUtil.GetPageFromQuerystring(ctx)

	authors, err := c.authorService.FindAll(page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(authors)
}

func (c *authorController) FindByName(ctx *fiber.Ctx) error {
	page := c.querystringUtil.GetPageFromQuerystring(ctx)

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
