package controllers

import (
	"net/http"
	"strconv"

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
}

type bookController struct {
	bookService   services.BookServiceInterface
	validatorUtil utils.ValidatorUtilInterface
}

func NewBookController(bookService services.BookServiceInterface, validatorUtil utils.ValidatorUtilInterface) BookControllerInterface {
	return &bookController{bookService, validatorUtil}
}

func (c *bookController) FindAll(ctx *fiber.Ctx) error {
	page := c.GetPageFromQuerystring(ctx)

	books, err := c.bookService.FindAll(page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(books)
}

func (c *bookController) FindByName(ctx *fiber.Ctx) error {
	page := c.GetPageFromQuerystring(ctx)

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
	bookFields := new(models.UpdateBookSchema)

	err := ctx.BodyParser(bookFields)
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

func (c *bookController) GetPageFromQuerystring(ctx *fiber.Ctx) int {
	result, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return 1
	}

	if result < 1 {
		result = 1
	}

	return result
}
