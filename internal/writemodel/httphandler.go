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
