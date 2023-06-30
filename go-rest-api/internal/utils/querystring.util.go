package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type QuerystringUtilInterface interface {
	GetPageFromQuerystring(ctx *fiber.Ctx) int
}

type queryStringUtil struct{}

func NewQuerystringUtil() QuerystringUtilInterface {
	return &queryStringUtil{}
}

func (q *queryStringUtil) GetPageFromQuerystring(ctx *fiber.Ctx) int {
	result, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return 1
	}

	if result < 1 {
		result = 1
	}

	return result
}
