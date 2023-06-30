package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RequestParamUtilInterface interface {
	GetPageFromQueryString(ctx *fiber.Ctx) int
	GetIntFromQueryString(ctx *fiber.Ctx, paramName string) int
}

type requestParamUtil struct{}

func NewRequestParamUtil() RequestParamUtilInterface {
	return &requestParamUtil{}
}

func (r *requestParamUtil) GetPageFromQueryString(ctx *fiber.Ctx) int {
	result, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return 1
	}

	if result < 1 {
		result = 1
	}

	return result
}

func (r *requestParamUtil) GetIntFromQueryString(ctx *fiber.Ctx, paramName string) int {
	queryStringInt, err := strconv.Atoi(ctx.Query(paramName))
	if err != nil || queryStringInt == 0 {
		return 0
	}

	return queryStringInt
}
