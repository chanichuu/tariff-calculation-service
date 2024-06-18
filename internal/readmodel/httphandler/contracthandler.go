//go:generate mockgen -source=contracthandler.go -destination=testing/contracthandler_mocks.go -package=testing ContractGetter, Validator

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

type ContractGetter interface {
	GetContracts(partitionId string) (*[]models.Contract, error)
	GetContract(partitionId, contractId string) (*models.Contract, error)
}

type ContractHandler struct {
	ContractRepo ContractGetter
	Validator    interfaces.Validator
}

func NewContractHandler() ContractHandler {
	return ContractHandler{
		ContractRepo: database.NewContractRepo(),
		Validator:    validation.NewValidator(),
	}
}

func (handler ContractHandler) HandleGetContracts(context *gin.Context) {
	pathParam := validation.PartitionId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParam); err != nil {
		return
	}

	contracts, err := handler.ContractRepo.GetContracts(pathParam.PartitionId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}

	context.JSON(http.StatusOK, contracts)
}

func (handler ContractHandler) HandleGetContract(context *gin.Context) {
	pathParams := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParams); err != nil {
		return
	}

	contract, err := handler.ContractRepo.GetContract(pathParams.PartitionId, pathParams.Id)
	if err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.JSON(http.StatusOK, contract)
}
