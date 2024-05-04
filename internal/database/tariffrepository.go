package database

import (
	"errors"
	"tariff-calculation-service/internal/models"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TariffRepo struct {
	DBClient
}

func NewTariffRepo() TariffRepo {
	return TariffRepo{
		DBClient: NewDBClient(),
	}
}

func (tr TariffRepo) GetKey(partitionId, tariffId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		tr.PartitionKey: &types.AttributeValueMemberS{Value: partitionId},
		tr.SortKey:      &types.AttributeValueMemberS{Value: TariffSortKeyPrefix + tariffId},
	}
}

func (tr TariffRepo) GetTariffs(partitionId string) (*[]models.Tariff, error) {
	tariffEntities, err := QueryEntities[models.Tariff](tr.DBClient, partitionId, TariffSortKeyPrefix)
	if err != nil {
		return nil, errors.New("failed to query tariffs")
	}
	tariffs := []models.Tariff{}
	for _, tariff := range tariffEntities {
		tariffs = append(tariffs, tariff.Data)
	}

	return &tariffs, nil
}

func (tr TariffRepo) GetTariff(partitionId, tariffId string) (*models.Tariff, error) {
	tariff, err := GetEntity[models.Tariff](tr.DBClient, tr.GetKey(partitionId, tariffId))
	if err != nil || tariff == nil {
		return &models.Tariff{}, err
	}

	return tariff, nil
}

func (tr TariffRepo) CreateTariff(partitionId string, tariff models.Tariff) (*models.Tariff, error) {
	return &tariff, nil
}

func (tr TariffRepo) UpdateTariff(partitionId string, tariff models.Tariff) error {
	return nil
}
func (tr TariffRepo) DeleteTariff(partitionId, tariffId string) error {
	return nil
}
