package database

import (
	"errors"
	dbtesting "tariff-calculation-service/internal/database/testing"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test/data"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testcase struct {
	Name             string
	PartitionId      string
	ContractId       string
	Mock             []func()
	expectedResponse any
}

func Test_GetContracts(t *testing.T) {
	// arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	contractRepo := ContractRepo{
		DBClient: testDBClient,
	}

	testcases := []testcase{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ContractId:  data.TestContractId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(data.TestContractQueryOutputWithoutTariffs, nil)
				},
			},
			expectedResponse: &data.Contracts,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			ContractId:  data.TestContractId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: []models.Contract{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualContracts, err := contractRepo.GetContracts(tc.PartitionId)
			// assert
			if err != nil {
				assert.Contains(t, "failed to query contracts", err.Error())
				assert.Nil(t, actualContracts)
			} else {
				assert.NotNil(t, actualContracts)
				assert.GreaterOrEqual(t, 1, len(*actualContracts))
				assert.Equal(t, tc.expectedResponse, actualContracts)
			}
		})
	}
}

func Test_GetContract(t *testing.T) {
	// arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	contractRepo := ContractRepo{
		DBClient: testDBClient,
	}

	testcases := []testcase{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ContractId:  data.TestContractId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(data.TestGetItemOutputContractWithoutTariffs, nil)
				},
			},
			expectedResponse: &data.Contract,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			ContractId:  data.TestContractId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: &models.Contract{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualContract, err := contractRepo.GetContract(tc.PartitionId, tc.ContractId)
			// assert
			if err != nil {
				assert.Contains(t, constants.ResourceNotFound, err.Error())
			} else {
				assert.NotNil(t, actualContract)
			}
			assert.Equal(t, tc.expectedResponse, actualContract)
		})
	}
}
