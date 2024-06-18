package validation

import (
	"net/http"
	"tariff-calculation-service/internal/models"

	"github.com/gin-gonic/gin"
)

type Validator struct {
}

func NewValidator() Validator {
	return Validator{}
}

func (Validator Validator) ValidateAndSetPathParams(ctx *gin.Context, objectPtr any) error {
	err := ctx.ShouldBindUri(objectPtr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return err
	}

	return nil
}
