package httphandler

import (
	"net/http"
	"strings"

	"tariff-calculation-service/internal/database"

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
		if strings.Contains(err.Error(), "ResourceNotFound") {
			context.IndentedJSON(http.StatusInternalServerError, nil)
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	context.IndentedJSON(http.StatusOK, provider)
}
