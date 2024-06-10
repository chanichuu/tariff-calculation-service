//go:generate mockgen -source=tariffhandler.go -destination=testing/tariffhandler_mocks.go -package=testing TariffGetter

package httphandler

import (
	"net/http"
	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg"

	"github.com/gin-gonic/gin"
)

type TariffGetter interface {
	GetTariffs(partitionId string) (*[]models.Tariff, error)
	GetTariff(partitionId, tariffId string) (*models.Tariff, error)
}

type TariffHandler struct {
	TariffRepo TariffGetter
}

func NewTariffHandler() TariffHandler {
	return TariffHandler{
		TariffRepo: database.NewTariffRepo(),
	}
}

func (handler TariffHandler) HandleGetTariffs(context *gin.Context) {
	partitionId := context.Param("pid")
	tariffs, err := handler.TariffRepo.GetTariffs(partitionId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}
	context.IndentedJSON(http.StatusOK, tariffs)
}

func (handler TariffHandler) HandleGetTariff(context *gin.Context) {
	partitionId := context.Param("pid")
	tariffId := context.Param("tid")
	tariff, err := handler.TariffRepo.GetTariff(partitionId, tariffId)
	if err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.IndentedJSON(http.StatusOK, tariff)
}
