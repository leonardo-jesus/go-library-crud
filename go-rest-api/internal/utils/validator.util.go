package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidatorUtilInterface interface {
	ValidateStruct(anyStruct interface{}) (errors []*ErrorResponse)
}

type validatorUtil struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

func NewValidatorUtil(validator *validator.Validate) ValidatorUtilInterface {
	return &validatorUtil{validator}
}

func (v *validatorUtil) ValidateStruct(anyStruct interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	err := v.validator.Struct(anyStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
