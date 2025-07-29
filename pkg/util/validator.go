package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s any) *AppError {
	if err := validate.Struct(s); err != nil {
		var errorList []FieldError
		for _, err := range err.(validator.ValidationErrors) {
			errorList = append(errorList, FieldError{
				Field: err.Field(),
				Error: fmt.Sprintf("%s is invalid (%s)", err.Field(), err.ActualTag()),
			})
		}
		return &AppError{
			Code:    "VALIDATION_ERROR",
			Message: "validation failed",
			Details: errorList,
		}
	}
	return nil
}
