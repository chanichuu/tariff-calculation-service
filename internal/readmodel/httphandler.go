package readmodel

import "tariff-calculation-service/internal/models"

func HandleGetTariffs(partitionId string) *[]models.Tariff {
	tariffs := []models.Tariff{}
	return &tariffs
}

func HandleGetTariff(partitionId, tariffId string) *models.Tariff {
	tariff := models.Tariff{}
	return &tariff
}
