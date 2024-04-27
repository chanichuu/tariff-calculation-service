package readmodel

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleGetHealth(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Service is healthy.")
}

func HandleGetVersion(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, os.Getenv("VERSION"))
}

func HandleGetRestVersion(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, os.Getenv("REST_API_VERSION"))
}

func HandleGetTariffs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func HandleGetTariff(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func HandleGetContracts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func HandleGetContract(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func HandleGetProviders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func HandleGetProvider(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}
