package data

import (
	"tariff-calculation-service/internal/models"
)

var Tariff = models.Tariff{
	Id:            TestTariffId,
	Name:          TestTariffName,
	Currency:      TestCurrency,
	ValidFrom:     TestValidFrom,
	ValidTo:       TestValidTo,
	TariffType:    TestTariffType,
	FixedTariff:   fixedTariff,
	DynamicTariff: dynamicTariff,
}

var Tariffs = []models.Tariff{Tariff}

var fixedTariff = models.FixedTariff{
	PricePerUnit: 64.5,
}

var dynamicTariff = models.DynamicTariff{
	HourlyTariffs: &[]models.HourlyTariff{hourlyTariff},
}

var hourlyTariff = models.HourlyTariff{
	StartTime:    TestValidFrom,
	ValidDays:    []uint8{0, 1, 2},
	PricePerUnit: 54.2,
}
