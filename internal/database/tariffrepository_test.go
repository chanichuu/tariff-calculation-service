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

type testcaseTariffRepo struct {
	Name             string
	PartitionId      string
	TariffId         string
	Mock             []func()
	expectedResponse any
}

func Test_GetTariffs(t *testing.T) {
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

	tariffRepo := TariffRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseTariffRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(data.TestGetQueryOutputTariff, nil)
				},
			},
			expectedResponse: &data.Tariffs,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: &models.Tariff{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualTariffs, err := tariffRepo.GetTariffs(tc.PartitionId)
			// assert
			if err != nil {
				assert.Contains(t, "failed to query tariffs", err.Error())
				assert.Nil(t, actualTariffs)
			} else {
				assert.NotNil(t, actualTariffs)
				assert.GreaterOrEqual(t, 1, len(*actualTariffs))
				assert.Equal(t, tc.expectedResponse, actualTariffs)
			}
		})
	}
}

func Test_GetTariff(t *testing.T) {
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

	tariffRepo := TariffRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseTariffRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(data.TestGetItemOutputTariff, nil)
				},
			},
			expectedResponse: &data.Tariff,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: &models.Tariff{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualTariff, err := tariffRepo.GetTariff(tc.PartitionId, tc.TariffId)
			// assert
			if err != nil {
				assert.Contains(t, constants.ResourceNotFound, err.Error())
			} else {
				assert.NotNil(t, actualTariff)
			}
			assert.Equal(t, tc.expectedResponse, actualTariff)
		})
	}
}

func Test_CreateTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	tariffRepo := TariffRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseTariffRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(data.TestPutItemOutputTariff, nil)
				},
			},
			expectedResponse: &data.Tariff,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, errors.New(constants.InternalServerError))
				},
			},
			expectedResponse: &models.Tariff{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualTariffPtr, err := tariffRepo.CreateTariff(tc.PartitionId, data.Tariff)
			// assert
			if err != nil {
				assert.Contains(t, constants.InternalServerError, err.Error())
			} else {
				assert.NotNil(t, actualTariffPtr)
			}
			assert.Equal(t, tc.expectedResponse, actualTariffPtr)
		})
	}
}

func Test_UpdateTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	tariffRepo := TariffRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseTariffRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().UpdateItem(gomock.Any(), gomock.Any()).Return(data.TestUpdateItemOutputTariff, nil)
				},
			},
			expectedResponse: nil,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().UpdateItem(gomock.Any(), gomock.Any()).Return(&dynamodb.UpdateItemOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: errors.New(constants.ResourceNotFound),
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			err := tariffRepo.UpdateTariff(tc.PartitionId, data.Tariff)
			// assert
			assert.Equal(t, tc.expectedResponse, err)
		})
	}
}

func Test_DeleteTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	tariffRepo := TariffRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseTariffRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().DeleteItem(gomock.Any(), gomock.Any()).Return(&dynamodb.DeleteItemOutput{}, nil)
				},
			},
			expectedResponse: nil,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			TariffId:    data.TestTariffId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().DeleteItem(gomock.Any(), gomock.Any()).Return(&dynamodb.DeleteItemOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: errors.New(constants.ResourceNotFound),
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			err := tariffRepo.DeleteTariff(tc.PartitionId, tc.TariffId)
			// assert
			if err != nil {
				assert.Contains(t, constants.ResourceNotFound, err.Error())
			}
			assert.Equal(t, tc.expectedResponse, err)
		})
	}
}
