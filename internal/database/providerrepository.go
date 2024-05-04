package database

import (
	"errors"
	"tariff-calculation-service/internal/models"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ProviderRepo struct {
	DBClient
}

func NewProviderRepo() ProviderRepo {
	return ProviderRepo{
		DBClient: NewDBClient(),
	}
}

func (pr ProviderRepo) GetKey(partitionId, providerId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		pr.PartitionKey: &types.AttributeValueMemberS{Value: partitionId},
		pr.SortKey:      &types.AttributeValueMemberS{Value: ProviderSortKeyPrefix + providerId},
	}
}

func (pr ProviderRepo) GetProviders(partitionId string) (*[]models.Provider, error) {
	providerEntities, err := QueryEntities[models.Provider](pr.DBClient, partitionId, ProviderSortKeyPrefix)
	if err != nil {
		return nil, errors.New("failed to query providers")
	}
	providers := []models.Provider{}

	for _, entity := range providerEntities {
		providers = append(providers, entity.Data)
	}

	return &providers, nil
}

func (pr ProviderRepo) GetProvider(partitionId, providerId string) (*models.Provider, error) {
	provider, err := GetEntity[models.Provider](pr.DBClient, pr.GetKey(partitionId, providerId))
	if err != nil || provider == nil {
		return &models.Provider{}, err
	}
	return provider, nil
}

func (pr ProviderRepo) CreateProvider(partitionId string, provider models.Provider) (*models.Provider, error) {
	return &provider, nil
}

func (pr ProviderRepo) UpdateProvider(partitionId string, provider models.Provider) error {
	return nil
}
func (pr ProviderRepo) DeleteProvider(partitionId, providerId string) error {
	return nil
}
