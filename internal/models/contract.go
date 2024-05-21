package models

type Contract struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Provider    string   `json:"provider"`
	Tariffs     []Tariff `json:"tariffs"`
}
