package database

import (
	"errors"
	"fmt"
	"tariff-calculation-service/internal/models"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ContractRepo struct {
	DBClient
}

func NewContractRepo() ContractRepo {
	return ContractRepo{
		NewDBClient(),
	}
}

func (cr ContractRepo) GetKey(partitionId, contractId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		cr.PartitionKey: &types.AttributeValueMemberS{Value: partitionId},
		cr.SortKey:      &types.AttributeValueMemberS{Value: ContractSortKeyPrefix + contractId},
	}
}

func (cr ContractRepo) GetContracts(partitionId string) (*[]models.Contract, error) {
	contractEntities, err := QueryEntities[models.Contract](cr.DBClient, partitionId, ContractSortKeyPrefix)
	if err != nil {
		return nil, errors.New("failed to query contracts")
	}
	contracts := []models.Contract{}

	for _, entity := range contractEntities {
		contracts = append(contracts, entity.Data)
	}

	return &contracts, nil
}

func (cr ContractRepo) GetContract(partitionId, contractId string) (*models.Contract, error) {
	contract, err := GetEntity[models.Contract](cr.DBClient, cr.GetKey(partitionId, contractId))
	if err != nil || contract == nil {
		return &models.Contract{}, err
	}
	return contract, nil
}

func (cr ContractRepo) CreateContract(partitionId string, contract models.Contract) (*models.Contract, error) {
	contractDB := DBEntity[models.Contract]{
		PartitionKey: partitionId,
		SortKey:      ContractSortKeyPrefix + contract.Id,
		Data:         contract,
	}
	err := PutEntity[DBEntity[models.Contract]](cr.DBClient, contractDB)
	if err != nil {
		return &models.Contract{}, err
	}

	return &contract, nil
}

func (cr ContractRepo) UpdateContract(partitionId string, contract models.Contract) error {
	dbUpdate := expression.Set(expression.Name("Data"), expression.Value(contract))
	expr, err := expression.NewBuilder().WithUpdate(dbUpdate).Build()
	if err != nil {
		return fmt.Errorf("error building expression %v", err)
	}
	err = UpdateEntity(cr.DBClient, cr.GetKey(partitionId, contract.Id), expr)

	return err
}

func (cr ContractRepo) DeleteContract(partitionId, contractId string) error {
	err := DeleteEntity(cr.DBClient, cr.GetKey(partitionId, contractId))

	return err
}
