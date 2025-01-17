package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/services"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/utils"
)

const NO_AUTHORS_FOUND_ERROR_MESSAGE = "no authors found"

type AuthorControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByName(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type authorController struct {
	authorService services.AuthorServiceInterface
	pageParamUtil utils.PageParamUtilInterface
}

func NewAuthorController(authorService services.AuthorServiceInterface, pageParamUtil utils.PageParamUtilInterface) AuthorControllerInterface {
	return &authorController{authorService, pageParamUtil}
}

func (c *authorController) FindAll(ctx *fiber.Ctx) error {
	page := c.pageParamUtil.ConvertPageStringToInt(ctx.Query("page"))

	authors, err := c.authorService.FindAll(page)
	if err != nil {
		if err.Error() == NO_AUTHORS_FOUND_ERROR_MESSAGE {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}

		log.Print(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(authors)
}

func (c *authorController) FindByName(ctx *fiber.Ctx) error {
	page := c.pageParamUtil.ConvertPageStringToInt(ctx.Query("page"))

	authors, err := c.authorService.FindByName(ctx.Query("name"), page)
	if err != nil {
		if err.Error() == NO_AUTHORS_FOUND_ERROR_MESSAGE {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}

		log.Print(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(authors)
}

func (c *authorController) Create(ctx *fiber.Ctx) error {
	err := c.authorService.Create()
	if err != nil {
		log.Print(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"success": "true"})
}
