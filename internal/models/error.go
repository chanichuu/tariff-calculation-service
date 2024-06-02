package models

import "tariff-calculation-service/pkg/constants"

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
