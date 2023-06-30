package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/services"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/utils"
)

type BookControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByName(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type bookController struct {
	bookService     services.BookServiceInterface
	validatorUtil   utils.ValidatorUtilInterface
	querystringUtil utils.QuerystringUtilInterface
}

func NewBookController(bookService services.BookServiceInterface, validatorUtil utils.ValidatorUtilInterface, querystringUtil utils.QuerystringUtilInterface) BookControllerInterface {
	return &bookController{bookService, validatorUtil, querystringUtil}
}

func (c *bookController) FindAll(ctx *fiber.Ctx) error {
	page := c.querystringUtil.GetPageFromQuerystring(ctx)

	books, err := c.bookService.FindAll(page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(books)
}

func (c *bookController) FindByName(ctx *fiber.Ctx) error {
	page := c.querystringUtil.GetPageFromQuerystring(ctx)

	books, err := c.bookService.FindByName(ctx.Query("name"), page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(books)
}

func (c *bookController) Create(ctx *fiber.Ctx) error {
	book := new(models.CreateBookSchema)

	err := ctx.BodyParser(book)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	errors := c.validatorUtil.ValidateStruct(book)
	if errors != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(errors)
	}

	err = c.bookService.Create(book)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"success": "true"})
}

func (c *bookController) Update(ctx *fiber.Ctx) error {
	bookIdParam, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	bookFields := new(models.UpdateBookSchema)
	bookFields.Id = bookIdParam

	err = ctx.BodyParser(bookFields)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	errors := c.validatorUtil.ValidateStruct(bookFields)
	if errors != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(errors)
	}

	err = c.bookService.Update(bookFields)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"success": "true"})
}

func (c *bookController) Delete(ctx *fiber.Ctx) error {
	bookIdParam, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.bookService.Delete(bookIdParam)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"success": "true"})
}
