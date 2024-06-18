//go:generate mockgen -source=tariffwritehandler.go -destination=testing/tariffwritehandler_mocks.go -package=testing TariffWriter

package writehandlers

import (
	"net/http"
	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/internal/interfaces"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg"
	"tariff-calculation-service/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TariffWriter interface {
	CreateTariff(partitionId string, tariff models.Tariff) (*models.Tariff, error)
	UpdateTariff(partitionId string, tariff models.Tariff) error
	DeleteTariff(partitionId, tariffId string) error
}

type TariffHandler struct {
	TariffWriter TariffWriter
	Validator    interfaces.Validator
}

func NewTariffHandler() TariffHandler {
	return TariffHandler{TariffWriter: database.NewTariffRepo(), Validator: validation.NewValidator()}
}

func (handler TariffHandler) HandlePostTariff(context *gin.Context) {
	pathParams := validation.PartitionId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	newTariff := models.Tariff{}
	if err := context.ShouldBindJSON(&newTariff); err != nil {
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return
	}

	newTariff.Id = uuid.New().String()

	tariff, err := handler.TariffWriter.CreateTariff(pathParams.PartitionId, newTariff)
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}

	context.JSON(http.StatusCreated, tariff)
}

func (handler TariffHandler) HandlePutTariff(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	tariff := models.Tariff{}
	if err := context.ShouldBindJSON(&tariff); err != nil {
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return
	}

	if tariff.Id == "" {
		tariff.Id = pathParams.Id
	}

	if err := handler.TariffWriter.UpdateTariff(pathParams.PartitionId, tariff); err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (handler TariffHandler) HandleDeleteTariff(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	if err := handler.TariffWriter.DeleteTariff(pathParams.PartitionId, pathParams.Id); err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
