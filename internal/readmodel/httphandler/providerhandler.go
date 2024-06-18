//go:generate mockgen -source=providerhandler.go -destination=testing/providerhandler_mocks.go -package=testing ProviderGetter

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

type ProviderGetter interface {
	GetProviders(partitionId string) (*[]models.Provider, error)
	GetProvider(partitionId, providerId string) (*models.Provider, error)
}

type ProviderHandler struct {
	ProviderRepo ProviderGetter
	Validator    interfaces.Validator
}

func NewProviderHandler() ProviderHandler {
	return ProviderHandler{
		ProviderRepo: database.NewProviderRepo(),
		Validator:    validation.NewValidator(),
	}
}

func (handler ProviderHandler) HandleGetProviders(context *gin.Context) {
	pathParam := validation.PartitionId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParam); err != nil {
		return
	}

	providers, err := handler.ProviderRepo.GetProviders(pathParam.PartitionId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}

	context.IndentedJSON(http.StatusOK, providers)
}

func (handler ProviderHandler) HandleGetProvider(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParams); err != nil {
		return
	}

	provider, err := handler.ProviderRepo.GetProvider(pathParams.PartitionId, pathParams.Id)
	if err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, provider)
}
