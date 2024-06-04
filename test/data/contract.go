package data

import (
	"tariff-calculation-service/internal/models"
)

var Contract = models.Contract{
	Id:          TestContractId,
	Name:        TestContractName,
	Description: TestContractDescription,
	StartDate:   TestContractStartDate,
	EndDate:     TestContractEndDate,
	Provider:    TestProviderId,
	Tariffs:     []models.Tariff{}, // todo change this to be an array of tariff ids
}

var Contracts = []models.Contract{
	Contract,
}
