package readmodel

import (
	"tariff-calculation-service/pkg/constants"

	"github.com/gin-gonic/gin"
)

func RouteReadmodelCalls(router *gin.Engine) {
	subRouter := router.Group(constants.BasePath)

	// Base routes
	subRouter.GET(constants.HealthPath, HandleGetHealth)
	subRouter.GET(constants.VersionPath, HandleGetVersion)
	subRouter.GET(constants.RestVersionPath, HandleGetRestVersion)

	// Tariff routes
	subRouter.GET(constants.TariffsPath, HandleGetTariffs)
	subRouter.GET(constants.SingleTariffPath, HandleGetTariff)
}
