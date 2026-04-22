package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) []ValidationError {
	var errors []ValidationError

	if valErrs, ok := err.(validator.ValidationErrors); ok {
		for _, f := range valErrs {
			errors = append(errors, ValidationError{
				Field:   strings.ToLower(f.Field()),
				Message: getErrorMessage(f),
			})
		}
	} else {
		errors = append(errors, ValidationError{
			Field:   "global",
			Message: "Invalid request payload",
		})
	}

	return errors
}

func getErrorMessage(f validator.FieldError) string {
	switch f.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required", f.Field())
	case "email":
		return "Please provide a valid email address"
	case "min":
		return fmt.Sprintf("The %s must be at least %s characters long", f.Field(), f.Param())
	case "max":
		return fmt.Sprintf("The %s must not exceed %s characters", f.Field(), f.Param())
	case "unique":
		return fmt.Sprintf("This %s is already taken", f.Field())
	default:
		return fmt.Sprintf("The %s field is invalid", f.Field())
	}
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func ErrorResponse(message string, errors interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Errors:  errors,
	}
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
