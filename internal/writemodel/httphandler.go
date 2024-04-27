package writemodel

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePostTariffs(context *gin.Context) {
	context.IndentedJSON(http.StatusCreated, nil)
}

func HandlePutTariff(context *gin.Context) {
	context.IndentedJSON(http.StatusNoContent, nil)
}

func HandleDeleteTariff(context *gin.Context) {
	context.IndentedJSON(http.StatusNoContent, nil)
}

func HandlePostContract(context *gin.Context) {
	context.IndentedJSON(http.StatusCreated, nil)
}

func HandlePutContract(context *gin.Context) {
	context.IndentedJSON(http.StatusNoContent, nil)
}

func HandleDeleteContract(context *gin.Context) {
	context.IndentedJSON(http.StatusNoContent, nil)
}

func HandlePostProvider(context *gin.Context) {
	context.IndentedJSON(http.StatusCreated, nil)
}

func HandlePutProvider(context *gin.Context) {
	context.IndentedJSON(http.StatusNoContent, nil)
}

func HandleDeleteProvider(context *gin.Context) {
	context.IndentedJSON(http.StatusNoContent, nil)
}
