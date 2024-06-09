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

type testcaseProviderRepo struct {
	Name             string
	PartitionId      string
	ProviderId       string
	Mock             []func()
	expectedResponse any
}

func Test_GetProviders(t *testing.T) {
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

	providerRepo := ProviderRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseProviderRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(data.TestProviderQueryOutput, nil)
				},
			},
			expectedResponse: &data.Providers,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
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
			actualProviders, err := providerRepo.GetProviders(tc.PartitionId)
			// assert
			if err != nil {
				assert.Contains(t, "failed to query providers", err.Error())
				assert.Nil(t, actualProviders)
			} else {
				assert.NotNil(t, actualProviders)
				assert.GreaterOrEqual(t, 1, len(*actualProviders))
				assert.Equal(t, tc.expectedResponse, actualProviders)
			}
		})
	}
}

func Test_GetProvider(t *testing.T) {
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

	providerRepo := ProviderRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseProviderRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(data.TestGetItemOutputProvider, nil)
				},
			},
			expectedResponse: &data.Provider,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{}, errors.New(constants.ResourceNotFound))
				},
			},
			expectedResponse: &models.Provider{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualProvider, err := providerRepo.GetProvider(tc.PartitionId, tc.ProviderId)
			// assert
			if err != nil {
				assert.Contains(t, constants.ResourceNotFound, err.Error())
			} else {
				assert.NotNil(t, actualProvider)
			}
			assert.Equal(t, tc.expectedResponse, actualProvider)
		})
	}
}

func Test_CreateProvider(t *testing.T) {
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

	providerRepo := ProviderRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseProviderRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(data.TestPutItemOutputProvider, nil)
				},
			},
			expectedResponse: &data.Provider,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, errors.New(constants.InternalServerError))
				},
			},
			expectedResponse: &models.Provider{},
		},
	}
	// act
	for _, tc := range testcases {
		for idx := range tc.Mock {
			tc.Mock[idx]()
		}
		t.Run(tc.Name, func(t *testing.T) {
			actualContract, err := providerRepo.CreateProvider(tc.PartitionId, data.Provider)
			// assert
			if err != nil {
				assert.Contains(t, constants.InternalServerError, err.Error())
			} else {
				assert.NotNil(t, actualContract)
			}
			assert.Equal(t, tc.expectedResponse, actualContract)
		})
	}
}

func Test_UpdateProvider(t *testing.T) {
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

	providerRepo := ProviderRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseProviderRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
			Mock: []func(){
				func() {
					mockDBManager.EXPECT().UpdateItem(gomock.Any(), gomock.Any()).Return(data.TestUpdateItemOutputProvider, nil)
				},
			},
			expectedResponse: nil,
		},
		{
			Name:        "Negative Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
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
			err := providerRepo.UpdateProvider(tc.PartitionId, data.Provider)
			// assert
			assert.Equal(t, tc.expectedResponse, err)
		})
	}
}

func Test_DeleteProvider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	providerRepo := ProviderRepo{
		DBClient: testDBClient,
	}

	testcases := []testcaseProviderRepo{
		{
			Name:        "Positive Test",
			PartitionId: data.TestPartitionId,
			ProviderId:  data.TestProviderId,
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
			ProviderId:  data.TestProviderId,
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
			err := providerRepo.DeleteProvider(tc.PartitionId, tc.ProviderId)
			// assert
			assert.Equal(t, tc.expectedResponse, err)
		})
	}
}
