package pkg

import (
	"net/http"
	"strings"

	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg/constants"

	"github.com/gin-gonic/gin"
)

func HandleResourceNotFoundAndInternalServerError(ctx *gin.Context, err error) {
	if strings.Contains(err.Error(), constants.ResourceNotFound) {
		ctx.IndentedJSON(http.StatusNotFound, models.NewResourceNotFoundError())
		return
	}
	ctx.IndentedJSON(http.StatusInternalServerError, models.NewInternalServerError())
}
