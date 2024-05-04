package httphandler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
}

func NewHttpHandler() HttpHandler {
	return HttpHandler{}
}

func (httpHandler HttpHandler) HandleGetHealth(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Service is healthy.")
}

func (httpHandler HttpHandler) HandleGetVersion(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, os.Getenv("VERSION"))
}

func (httpHandler HttpHandler) HandleGetRestVersion(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, os.Getenv("REST_API_VERSION"))
}
