package database

import (
	"tariff-calculation-service/internal/models"
)

func GetProviders(partitionId string) *[]models.Provider {
	providers := []models.Provider{}
	return &providers
}

func GetProvider(partitionId, providerId string) *models.Provider {
	provider := models.Provider{}
	return &provider
}

func CreateProvider(partitionId string, provider models.Provider) (*models.Provider, error) {
	return &provider, nil
}

func UpdateProvider(partitionId string, provider models.Provider) error {
	return nil
}
func DeleteProvider(partitionId, providerId string) error {
	return nil
}
