package models

import (
	"tariff-calculation-service/pkg/enums"
)

type Tariff struct {
	Id         string           `json:"id"`
	Name       string           `json:"name"`
	Currency   string           `json:"currency"`
	ValidFrom  string           `json:"validFrom"`
	ValidTo    string           `json:"validTo"`
	TariffType enums.TariffType `json:"tariffType"`
}

type FixedTariff struct {
	Tariff
	PricePerUnit float64 `json:"pricePerUnit"`
}

type DynamicTariff struct {
	Tariff
	HourlyTariffs *[]HourlyTariff `json:"hourlyTariffs"`
}

type HourlyTariff struct {
	StartTime    string           `json:"startTime"`
	ValidDays    []enums.WeekDays `json:"validDays"`
	PricePerUnit float64          `json:"pricePerUnit"`
}
