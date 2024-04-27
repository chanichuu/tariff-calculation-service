package database

import (
	"tariff-calculation-service/internal/models"
)

func GetTariffs(partitionId string) *[]models.Tariff {
	tariffs := []models.Tariff{}
	return &tariffs
}

func GetTariff(partitionId, tariffId string) *models.Tariff {
	tariff := models.Tariff{}
	return &tariff
}

func CreateTariff(partitionId string, tariff models.Tariff) (*models.Tariff, error) {
	return &tariff, nil
}

func UpdateTariff(partitionId string, tariff models.Tariff) error {
	return nil
}
func DeleteTariff(partitionId, tariffId string) error {
	return nil
}
