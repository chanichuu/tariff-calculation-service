package models

import "tariff-calculation-service/pkg/enums"

type Contract struct {
	Id          string
	Name        string
	Description string
	StartDate   string
	EndDate     string
	Provider    string
	Tariffs     []Tariff
	TariffType  enums.TariffType
}
