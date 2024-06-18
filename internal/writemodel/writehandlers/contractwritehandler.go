//go:generate mockgen -source=contractwritehandler.go -destination=testing/contractwritehandler_mocks.go -package=testing ContractWriter

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

type ContractWriter interface {
	CreateContract(partitionId string, contract models.Contract) (*models.Contract, error)
	UpdateContract(partitionId string, contract models.Contract) error
	DeleteContract(partitionId, contractId string) error
}

type ContractWriteHandler struct {
	ContractWriter ContractWriter
	Validator      interfaces.Validator
}

func NewContractWriteHandler() ContractWriteHandler {
	return ContractWriteHandler{ContractWriter: database.NewContractRepo(), Validator: validation.NewValidator()}
}

func (handler ContractWriteHandler) HandlePostContract(context *gin.Context) {
	pathParam := validation.PartitionId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParam); err != nil {
		return
	}

	newContract := models.Contract{}
	if err := context.ShouldBindJSON(&newContract); err != nil {
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return
	}

	newContract.Id = uuid.New().String()

	contract, err := handler.ContractWriter.CreateContract(pathParam.PartitionId, newContract)
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}

	context.JSON(http.StatusCreated, contract)
}

func (handler ContractWriteHandler) HandlePutContract(context *gin.Context) {
	pathParam := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParam); err != nil {
		return
	}

	contract := models.Contract{}
	if err := context.ShouldBindJSON(&contract); err != nil {
		context.JSON(http.StatusBadRequest, models.NewBadRequestFieldValidationError(err))
		return
	}

	if contract.Id == "" {
		contract.Id = pathParam.Id
	}

	if err := handler.ContractWriter.UpdateContract(pathParam.PartitionId, contract); err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (handler ContractWriteHandler) HandleDeleteContract(context *gin.Context) {
	pathParam := validation.PartitionIdWithId{}
	if err := handler.Validator.ValidateAndSetPathParams(context, pathParam); err != nil {
		return
	}

	if err := handler.ContractWriter.DeleteContract(pathParam.PartitionId, pathParam.Id); err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
