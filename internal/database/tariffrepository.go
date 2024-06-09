package database

import (
	"errors"
	"fmt"
	"tariff-calculation-service/internal/models"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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
	tariffDB := DBEntity[models.Tariff]{
		PartitionKey: partitionId,
		SortKey:      TariffSortKeyPrefix + tariff.Id,
		Data:         tariff,
	}
	err := PutEntity[DBEntity[models.Tariff]](tr.DBClient, tariffDB)
	if err != nil {
		return &models.Tariff{}, err
	}
	return &tariff, nil
}

func (tr TariffRepo) UpdateTariff(partitionId string, tariff models.Tariff) error {
	dbUpdate := expression.Set(expression.Name("Data"), expression.Value(tariff))
	expr, err := expression.NewBuilder().WithUpdate(dbUpdate).Build()
	if err != nil {
		return fmt.Errorf("error building expression %v", err)
	}
	err = UpdateEntity(tr.DBClient, tr.GetKey(partitionId, tariff.Id), expr)

	return err
}

func (tr TariffRepo) DeleteTariff(partitionId, tariffId string) error {
	err := DeleteEntity(tr.DBClient, tr.GetKey(partitionId, tariffId))

	return err
}
