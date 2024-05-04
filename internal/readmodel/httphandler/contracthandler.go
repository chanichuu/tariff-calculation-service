package httphandler

import (
	"net/http"
	"strings"

	"tariff-calculation-service/internal/database"

	"github.com/gin-gonic/gin"
)

type ContractHandler struct {
	ContractRepo database.ContractRepo
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
		context.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, contracts)
}

func (handler ContractHandler) HandleGetContract(context *gin.Context) {
	partitionId := context.Param("pid")
	contractId := context.Param("cid")
	contract, err := handler.ContractRepo.GetContract(partitionId, contractId)
	if err != nil {
		if strings.Contains(err.Error(), "ResourceNotFound") {
			context.IndentedJSON(http.StatusNotFound, nil)
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, contract)
}
