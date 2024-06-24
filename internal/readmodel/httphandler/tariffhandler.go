//go:generate mockgen -source=tariffhandler.go -destination=testing/tariffhandler_mocks.go -package=testing TariffGetter

package httphandler

import (
	"net/http"
	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/internal/interfaces"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg"
	"tariff-calculation-service/pkg/validation"

	"github.com/gin-gonic/gin"
)

type TariffGetter interface {
	GetTariffs(partitionId string) (*[]models.Tariff, error)
	GetTariff(partitionId, tariffId string) (*models.Tariff, error)
}

type TariffHandler struct {
	TariffRepo TariffGetter
	Validator  interfaces.Validator
}

func NewTariffHandler() TariffHandler {
	return TariffHandler{
		TariffRepo: database.NewTariffRepo(),
		Validator:  validation.NewValidator(),
	}
}

func (handler TariffHandler) HandleGetTariffs(context *gin.Context) {
	pathParam := validation.PartitionId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParam); err != nil {
		return
	}

	tariffs, err := handler.TariffRepo.GetTariffs(pathParam.PartitionId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}
	context.IndentedJSON(http.StatusOK, tariffs)
}

func (handler TariffHandler) HandleGetTariff(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	tariff, err := handler.TariffRepo.GetTariff(pathParams.PartitionId, pathParams.Id)
	if err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.IndentedJSON(http.StatusOK, tariff)
}
