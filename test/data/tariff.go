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
	HourlyTariffs: []models.HourlyTariff{hourlyTariff},
}

var hourlyTariff = models.HourlyTariff{
	StartTime:    TestValidFrom,
	ValidDays:    []uint8{0, 1, 2},
	PricePerUnit: 54.2,
}

var TariffInvalidHourlyStartTime = models.Tariff{
	Id:            TestTariffId,
	Name:          TestTariffName,
	Currency:      TestCurrency,
	ValidFrom:     TestValidFrom,
	ValidTo:       TestValidTo,
	TariffType:    TestTariffType,
	FixedTariff:   fixedTariff,
	DynamicTariff: dynamicTariffInvalid,
}

var dynamicTariffInvalid = models.DynamicTariff{
	HourlyTariffs: []models.HourlyTariff{hourlyTariffInvalidStartTime},
}

var hourlyTariffInvalidStartTime = models.HourlyTariff{
	StartTime:    "",
	ValidDays:    []uint8{0, 1, 2},
	PricePerUnit: 54.2,
}

var TariffInvalidHourlyValidDays = models.Tariff{
	Id:            TestTariffId,
	Name:          TestTariffName,
	Currency:      TestCurrency,
	ValidFrom:     TestValidFrom,
	ValidTo:       TestValidTo,
	TariffType:    TestTariffType,
	FixedTariff:   fixedTariff,
	DynamicTariff: dynamicTariffInvalidValidDays,
}

var dynamicTariffInvalidValidDays = models.DynamicTariff{
	HourlyTariffs: []models.HourlyTariff{hourlyTariffInvalidValidDays},
}

var hourlyTariffInvalidValidDays = models.HourlyTariff{
	StartTime:    TestValidFrom,
	ValidDays:    []uint8{9},
	PricePerUnit: 54.2,
}

var TariffInvalidHourlyPricePerUnit = models.Tariff{
	Id:            TestTariffId,
	Name:          TestTariffName,
	Currency:      TestCurrency,
	ValidFrom:     TestValidFrom,
	ValidTo:       TestValidTo,
	TariffType:    TestTariffType,
	FixedTariff:   fixedTariff,
	DynamicTariff: dynamicTariffInvalidPricePerUnit,
}

var dynamicTariffInvalidPricePerUnit = models.DynamicTariff{
	HourlyTariffs: []models.HourlyTariff{hourlyTariffInvalidPricePerUnit},
}

var hourlyTariffInvalidPricePerUnit = models.HourlyTariff{
	StartTime:    TestValidFrom,
	ValidDays:    []uint8{1},
	PricePerUnit: -1,
}
