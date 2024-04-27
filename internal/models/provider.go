package models

type Provider struct {
	Id      string
	Name    string
	Email   string
	Address Address
}

type Address struct {
	Street     string
	PostalCode uint16
	City       string
	Country    string
}
