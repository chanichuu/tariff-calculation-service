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

	// Contract routes
	subRouter.POST(constants.ContractsPath)
	subRouter.PUT(constants.SingleContractPath)
	subRouter.DELETE(constants.SingleContractPath)

	// Provider routes
	subRouter.POST(constants.ProvidersPath)
	subRouter.PUT(constants.SingleProviderPath)
	subRouter.DELETE(constants.SingleProviderPath)
}
