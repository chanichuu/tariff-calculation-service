package httphandler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ServiceHealth = "Service is healthy."
)

type HttpHandler struct {
}

func NewHttpHandler() HttpHandler {
	return HttpHandler{}
}

func (httpHandler HttpHandler) HandleGetHealth(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, ServiceHealth)
}

func (httpHandler HttpHandler) HandleGetVersion(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, os.Getenv("VERSION"))
}

func (httpHandler HttpHandler) HandleGetRestVersion(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, os.Getenv("REST_API_VERSION"))
}
