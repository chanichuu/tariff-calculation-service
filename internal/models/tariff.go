package models

import (
	"tariff-calculation-service/pkg/enums"
)

type Tariff struct {
	Id            string           `json:"id" binding:"uuid"`
	Name          string           `json:"name" binding:"required,max=64"`
	Currency      string           `json:"currency" binding:"required,iso4217"`
	ValidFrom     string           `json:"validFrom" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	ValidTo       string           `json:"validTo" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	TariffType    enums.TariffType `json:"tariffType" binding:"required,max=128"`
	FixedTariff   FixedTariff      `json:"fixedTariff"`
	DynamicTariff DynamicTariff    `json:"dynamicTariff"`
}

type FixedTariff struct {
	PricePerUnit float64 `json:"pricePerUnit" binding:"required,gte=0"`
}

type DynamicTariff struct {
	HourlyTariffs *[]HourlyTariff `json:"hourlyTariffs"`
}

type HourlyTariff struct {
	StartTime    string  `json:"startTime" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	ValidDays    []uint8 `json:"validDays" binding:"required,min=1,max=7,gte=1,lte=7"`
	PricePerUnit float64 `json:"pricePerUnit" binding:"required,gte=0"`
}
