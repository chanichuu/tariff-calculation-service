//go:generate mockgen -source=contracthandler.go -destination=testing/contracthandler_mocks.go -package=testing ContractGetter

package httphandler

import (
	"net/http"

	"tariff-calculation-service/internal/database"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg"

	"github.com/gin-gonic/gin"
)

type ContractGetter interface {
	GetContracts(partitionId string) (*[]models.Contract, error)
	GetContract(partitionId, contractId string) (*models.Contract, error)
}

type ContractHandler struct {
	ContractRepo ContractGetter
}

func NewContractHandler() ContractHandler {
	return ContractHandler{
		ContractRepo: database.NewContractRepo(),
	}
}

func (handler ContractHandler) HandleGetContracts(context *gin.Context) {
	partitionId := context.Param("pid")
	contracts, err := handler.ContractRepo.GetContracts(partitionId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, models.NewInternalServerError())
		return
	}

	context.IndentedJSON(http.StatusOK, contracts)
}

func (handler ContractHandler) HandleGetContract(context *gin.Context) {
	partitionId := context.Param("pid")
	contractId := context.Param("cid")
	contract, err := handler.ContractRepo.GetContract(partitionId, contractId)
	if err != nil {
		pkg.HandleResourceNotFoundAndInternalServerError(context, err)
		return
	}

	context.IndentedJSON(http.StatusOK, contract)
}
