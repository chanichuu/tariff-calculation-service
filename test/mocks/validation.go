package mocks

import (
	"errors"
	"net/http"
	"reflect"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/internal/readmodel/httphandler/testing"
	"tariff-calculation-service/pkg/validation"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func NewValidatorPathPositive(mockController *gomock.Controller) *testing.MockValidator {
	mockValidator := testing.NewMockValidator(mockController)
	mockValidator.EXPECT().ValidateAndSetPathParams(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(context *gin.Context, objPtr any) error {
		setByType(context, objPtr)
		return nil
	})
	return mockValidator
}

func NewValidatorPathNegative(mockController *gomock.Controller) *testing.MockValidator {
	mockValidatorPathNegative := testing.NewMockValidator(mockController)
	mockValidatorPathNegative.EXPECT().ValidateAndSetPathParams(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(context *gin.Context, objPtr any) error {
		err := errors.New("ValidationError")
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return err
	})

	return mockValidatorPathNegative
}

func setByType(context *gin.Context, objPtr any) {
	switch objPtr.(type) {
	case *validation.PartitionId:
		reflect.ValueOf(objPtr).Elem().Set(reflect.ValueOf(
			validation.PartitionId{
				PartitionId: context.Param("partitionId"),
			},
		))
	case *validation.PartitionIdWithId:
		reflect.ValueOf(objPtr).Elem().Set(reflect.ValueOf(
			validation.PartitionIdWithId{
				PartitionId: context.Param("partitionId"),
				Id:          context.Param("id"),
			},
		))
	}
}
