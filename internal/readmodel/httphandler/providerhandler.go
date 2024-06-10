//go:generate mockgen -source=providerhandler.go -destination=testing/providerhandler_mocks.go -package=testing ProviderGetter

package httphandler

import (
	"net/http"

	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg"

	"github.com/gin-gonic/gin"
)

type ProviderGetter interface {
	GetProviders(partitionId string) (*[]models.Provider, error)
	GetProvider(partitionId, providerId string) (*models.Provider, error)
}

type ProviderHandler struct {
	ProviderRepo ProviderGetter
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
		context.IndentedJSON(http.StatusInternalServerError, models.NewInternalServerError())
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
