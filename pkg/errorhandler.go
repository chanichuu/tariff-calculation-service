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
		ctx.JSON(http.StatusNotFound, models.NewResourceNotFoundError())
		return
	}
	ctx.JSON(http.StatusInternalServerError, models.NewInternalServerError())
}
