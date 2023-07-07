package utils

import (
	"strconv"
)

type PageParamUtilInterface interface {
	ConvertPageStringToInt(page string) int
}

type pageParamUtil struct{}

func NewPageParamUtil() PageParamUtilInterface {
	return &pageParamUtil{}
}

func (r *pageParamUtil) ConvertPageStringToInt(page string) int {
	result, err := strconv.Atoi(page)
	if err != nil || result < 1 {
		return 1
	}

	return result
}
