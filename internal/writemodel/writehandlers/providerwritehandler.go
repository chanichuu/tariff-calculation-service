//go:generate mockgen -source=providerwritehandler.go -destination=testing/providerwritehandler_mocks.go -package=testing ProviderWriter

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

type ProviderWriter interface {
	CreateProvider(partitionId string, provider models.Provider) (*models.Provider, error)
	UpdateProvider(partitionId string, provider models.Provider) error
	DeleteProvider(partitionId, providerId string) error
}

type ProviderHandler struct {
	ProviderWriter ProviderWriter
	Validator      interfaces.Validator
}

func NewProviderHandler() ProviderHandler {
	return ProviderHandler{ProviderWriter: database.NewProviderRepo(), Validator: validation.NewValidator()}
}

func (handler ProviderHandler) HandlePostProvider(context *gin.Context) {
	pathParams := validation.PartitionId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	newProvider := models.Provider{}
	if err := context.ShouldBindJSON(&newProvider); err != nil {
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return
	}

	newProvider.Id = uuid.New().String()

	provider, err := handler.ProviderWriter.CreateProvider(pathParams.PartitionId, newProvider)
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}

	context.JSON(http.StatusCreated, provider)
}

func (handler ProviderHandler) HandlePutProvider(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	provider := models.Provider{}
	if err := context.ShouldBindJSON(&provider); err != nil {
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return
	}

	if provider.Id == "" {
		provider.Id = pathParams.Id
	}

	if err := handler.ProviderWriter.UpdateProvider(pathParams.PartitionId, provider); err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (handler ProviderHandler) HandleDeleteProvider(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, &pathParams); err != nil {
		return
	}

	if err := handler.ProviderWriter.DeleteProvider(pathParams.PartitionId, pathParams.Id); err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
