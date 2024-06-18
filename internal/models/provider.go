package models

type Provider struct {
	Id      string  `json:"id" binding:"uuid"`
	Name    string  `json:"name" binding:"required,max=64"`
	Email   string  `json:"email" binding:"email"`
	Address Address `json:"address"`
}

type Address struct {
	Street      string `json:"street" binding:"required,max=64"`
	PostalCode  string `json:"postalCode" binding:"required,max=12"`
	City        string `json:"city" binding:"required,max=64"`
	CountryCode string `json:"country" binding:"required,iso3166_1_alpha3"`
}
