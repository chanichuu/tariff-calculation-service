package constants

const (
	BasePath           string = "/api/v1/partitions/:pid"
	HealthPath         string = "/health"
	VersionPath        string = "/version"
	RestVersionPath    string = "/rest-version"
	TariffsPath        string = "/tariffs"
	SingleTariffPath   string = TariffsPath + "/:tid"
	ContractsPath      string = "/contracts"
	SingleContractPath string = ContractsPath + "/:cid"
	ProvidersPath      string = "/providers"
	SingleProviderPath string = ProvidersPath + "/:id"
)
