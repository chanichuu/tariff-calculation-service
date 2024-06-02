package httphandler

import (
	"net/http"
	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/pkg"

	"github.com/gin-gonic/gin"
)

type TariffHandler struct {
	TariffRepo database.TariffRepo
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
		context.IndentedJSON(http.StatusInternalServerError, nil)
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
	}

	context.IndentedJSON(http.StatusOK, tariff)
}
