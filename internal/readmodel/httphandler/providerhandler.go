package httphandler

import (
	"net/http"

	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/pkg"

	"github.com/gin-gonic/gin"
)

type ProviderHandler struct {
	ProviderRepo database.ProviderRepo
}

func NewProviderHandler() ProviderHandler {
	return ProviderHandler{
		ProviderRepo: database.NewProviderRepo(),
	}
}

func (handler ProviderHandler) HandleGetProviders(context *gin.Context) {
	partitionId := context.Param("pid")
	providers, err := handler.ProviderRepo.GetProviders(partitionId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, providers)
}

func (handler ProviderHandler) HandleGetProvider(context *gin.Context) {
	partitionId := context.Param("pid")
	providerId := context.Param("id")
	provider, err := handler.ProviderRepo.GetProvider(partitionId, providerId)
	if err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, provider)
}
