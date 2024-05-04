package readmodel

import (
	"tariff-calculation-service/internal/readmodel/httphandler"
	"tariff-calculation-service/pkg/constants"

	"github.com/gin-gonic/gin"
)

func RouteReadmodelCalls(router *gin.Engine) {
	subRouter := router.Group(constants.BasePath)
	serviceHandler := httphandler.NewHttpHandler()
	tariffHandler := httphandler.NewTariffHandler()
	contractHandler := httphandler.NewContractHandler()
	providerHandler := httphandler.NewProviderHandler()

	// Base routes
	subRouter.GET(constants.HealthPath, serviceHandler.HandleGetHealth)
	subRouter.GET(constants.VersionPath, serviceHandler.HandleGetVersion)
	subRouter.GET(constants.RestVersionPath, serviceHandler.HandleGetRestVersion)

	// Tariff routes
	subRouter.GET(constants.TariffsPath, tariffHandler.HandleGetTariffs)
	subRouter.GET(constants.SingleTariffPath, tariffHandler.HandleGetTariff)

	// Contract routes
	subRouter.GET(constants.ContractsPath, contractHandler.HandleGetContracts)
	subRouter.GET(constants.SingleContractPath, contractHandler.HandleGetContract)

	// Provider routes
	subRouter.GET(constants.ProvidersPath, providerHandler.HandleGetProviders)
	subRouter.GET(constants.SingleProviderPath, providerHandler.HandleGetProvider)
}
