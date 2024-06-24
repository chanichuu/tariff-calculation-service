package models

import (
	"errors"
	"fmt"
	"strings"
	"tariff-calculation-service/pkg/constants"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Code   int
	Name   string
	Detail string
}

func NewResourceNotFoundError() Error {
	return Error{
		Code:   404,
		Name:   constants.ResourceNotFound,
		Detail: "Resource not found",
	}
}

func NewInternalServerError() Error {
	return Error{
		Code:   500,
		Name:   constants.InternalServerError,
		Detail: "Internal Server Error",
	}
}

func NewBadRequestFieldValidationError(err error) Error {
	var validationError validator.ValidationErrors
	if !errors.As(err, &validationError) {
		return NewBadRequestError(err)
	}
	var validationErrors []string
	for _, field := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, field.Field())
	}

	return Error{
		Code:   400,
		Name:   constants.BadRequest,
		Detail: fmt.Sprintf("Invalid value: %s", strings.Join(validationErrors, ", ")),
	}
}

func NewBadRequestError(err error) Error {
	return Error{
		Code:   400,
		Name:   constants.BadRequest,
		Detail: err.Error(),
	}
}
