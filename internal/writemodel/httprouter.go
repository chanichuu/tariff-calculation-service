package writemodel

import (
	"tariff-calculation-service/internal/writemodel/writehandlers"
	"tariff-calculation-service/pkg/constants"

	"github.com/gin-gonic/gin"
)

func RouteReadmodelCalls(router *gin.Engine) {
	subRouter := router.Group(constants.BasePath)
	contractHandler := writehandlers.NewContractWriteHandler()
	providerHandler := writehandlers.NewProviderHandler()
	tariffHandler := writehandlers.NewTariffHandler()

	// Tariff routes
	subRouter.POST(constants.TariffsPath, tariffHandler.HandlePostTariff)
	subRouter.PUT(constants.SingleTariffPath, tariffHandler.HandlePutTariff)
	subRouter.DELETE(constants.SingleTariffPath, tariffHandler.HandleDeleteTariff)

	// Contract routes
	subRouter.POST(constants.ContractsPath, contractHandler.HandlePostContract)
	subRouter.PUT(constants.SingleContractPath, contractHandler.HandlePutContract)
	subRouter.DELETE(constants.SingleContractPath, contractHandler.HandleDeleteContract)

	// Provider routes
	subRouter.POST(constants.ProvidersPath, providerHandler.HandlePostProvider)
	subRouter.PUT(constants.SingleProviderPath, providerHandler.HandlePutProvider)
	subRouter.DELETE(constants.SingleProviderPath, providerHandler.HandleDeleteProvider)
}
