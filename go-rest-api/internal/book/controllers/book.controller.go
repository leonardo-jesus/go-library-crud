package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/services"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/utils"
)

const NO_BOOKS_FOUND_ERROR_MESSAGE = "no books found"

type BookControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByFilteredBooks(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type bookController struct {
	bookService      services.BookServiceInterface
	validatorUtil    utils.ValidatorUtilInterface
	requestParamUtil utils.RequestParamUtilInterface
}

func NewBookController(bookService services.BookServiceInterface, validatorUtil utils.ValidatorUtilInterface, requestParamUtil utils.RequestParamUtilInterface) BookControllerInterface {
	return &bookController{bookService, validatorUtil, requestParamUtil}
}

func (c *bookController) FindAll(ctx *fiber.Ctx) error {
	page := c.requestParamUtil.GetPageFromQueryString(ctx)

	books, err := c.bookService.FindAll(page)
	if err != nil {
		if err.Error() == NO_BOOKS_FOUND_ERROR_MESSAGE {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(books)
}

func (c *bookController) FindByFilteredBooks(ctx *fiber.Ctx) error {
	page := c.requestParamUtil.GetPageFromQueryString(ctx)

	var nameValue *string
	var nameFromQuery = ctx.Query("name")
	if nameFromQuery != "" {
		nameValue = new(string)
		*nameValue = nameFromQuery
	}

	var authorSlice *[]int
	authorsFromQuery := ctx.Query("authors")
	if authorsFromQuery != "" {
		authorSlice = new([]int)
		authorIds := strings.Split(authorsFromQuery, ",")
		for _, idString := range authorIds {
			id, err := strconv.Atoi(idString)
			if err != nil {
				ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}
			*authorSlice = append(*authorSlice, id)
		}
	}

	var publicationYearValue, editionValue *int
	publicationYearFromQuery := c.requestParamUtil.GetIntFromQueryString(ctx, "publicationYear")
	if publicationYearFromQuery != 0 {
		publicationYearValue = new(int)
		*publicationYearValue = publicationYearFromQuery
	}

	editionFromQuery := c.requestParamUtil.GetIntFromQueryString(ctx, "edition")
	if editionFromQuery != 0 {
		editionValue = new(int)
		*editionValue = editionFromQuery
	}

	queryStringFilters := models.FilteredBookSchema{
		Name:            nameValue,
		PublicationYear: publicationYearValue,
		Edition:         editionValue,
		Authors:         authorSlice,
	}

	errors := c.validatorUtil.ValidateStruct(queryStringFilters)
	if errors != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(errors)
	}

	books, err := c.bookService.FindByFilteredBooks(queryStringFilters, page)
	if err != nil {
		if err.Error() == NO_BOOKS_FOUND_ERROR_MESSAGE {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

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
		if err.Error() == NO_BOOKS_FOUND_ERROR_MESSAGE {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

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
