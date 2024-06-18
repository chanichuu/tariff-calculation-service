package models

type Contract struct {
	Id          string   `json:"id" binding:"uuid"`
	Name        string   `json:"name" binding:"required,max=64"`
	Description string   `json:"description" binding:"max=128"`
	StartDate   string   `json:"startDate" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndDate     string   `json:"endDate" binding:"datetime=2006-01-02T15:04:05Z07:00"`
	Provider    string   `json:"provider"`
	Tariffs     []string `json:"tariffs" binding:"dive,uuid"`
}
