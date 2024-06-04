package data

import (
	"tariff-calculation-service/internal/models"
)

var Provider = models.Provider{
	Id:      TestProviderId,
	Name:    "TestProvider",
	Email:   "test@provider.com",
	Address: address,
}

var address = models.Address{
	Street:     "TestStreet",
	PostalCode: "107-6001",
	City:       "Tokyo",
	Country:    "Japan",
}

var Providers = []models.Provider{
	Provider,
}
