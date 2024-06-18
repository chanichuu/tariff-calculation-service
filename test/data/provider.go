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
	Street:      "TestStreet",
	PostalCode:  "107-6001",
	City:        "Tokyo",
	CountryCode: "JPN",
}

var Providers = []models.Provider{
	Provider,
}
