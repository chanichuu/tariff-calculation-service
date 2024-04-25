package writemodel

import (
	"tariff-calculation-service/pkg/constants"

	"github.com/gin-gonic/gin"
)

func RouteReadmodelCalls(router *gin.Engine) {
	subRouter := router.Group(constants.BasePath)

	// Tariff routes
	subRouter.POST(constants.TariffsPath)
	subRouter.PUT(constants.SingleTariffPath)
	subRouter.DELETE(constants.SingleTariffPath)
}
