package database

import (
	"tariff-calculation-service/internal/models"
)

func GetContracts(partitionId string) *[]models.Contract {
	contracts := []models.Contract{}
	return &contracts
}

func GetContract(partitionId, contractId string) *models.Contract {
	contract := models.Contract{}
	return &contract
}

func CreateContract(partitionId string, contract models.Contract) (*models.Contract, error) {
	return &contract, nil
}

func UpdateContract(partitionId string, contract models.Contract) error {
	return nil
}
func DeleteContract(partitionId, contractId string) error {
	return nil
}
