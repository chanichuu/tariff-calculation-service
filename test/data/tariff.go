package data

import (
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg/enums"
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
	ValidDays:    []enums.WeekDays{enums.MO, enums.TU, enums.WE},
	PricePerUnit: 54.2,
}
