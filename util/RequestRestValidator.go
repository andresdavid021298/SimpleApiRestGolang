package util

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RequestRestValidator struct {
	Validator *validator.Validate
}

func (validator *RequestRestValidator) Validate(i interface{}) error {
	if err := validator.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ValidateRequestErrors(bindError error, validateError error) string {
	if validateError != nil {
		return validateError.Error()
	}
	if bindError != nil {
		return bindError.Error()
	}
	return ""
}
