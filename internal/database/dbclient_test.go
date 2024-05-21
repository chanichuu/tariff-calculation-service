package database

import (
	"errors"
	dbtesting "tariff-calculation-service/internal/database/testing"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test/data"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	POSITIVE = "Positive"
	NEGATIVE = "Negative"
	CONTRACT = "Contract"
	TARIFF   = "Tariff"
	PROVIDER = "Provider"
)

type TestCase struct {
	Name string
	Type string
	Mock func()
}

var testKey = map[string]types.AttributeValue{
	"TestPartitionKey": &types.AttributeValueMemberS{Value: data.TestPartitionId},
	"TestSortKey":      &types.AttributeValueMemberS{Value: data.TestSortKey},
}

func TestUnit_GetEntity(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Positive Test Contract",
			Type: CONTRACT,
			Mock: func() {
				mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(data.TestGetItemOutputContract, nil)
			},
		},
		{
			Name: "Positive Test Tariff",
			Type: TARIFF,
			Mock: func() {
				mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(data.TestGetItemOutputTariff, nil)
			},
		},
		{
			Name: "Positive Test Provider",
			Type: PROVIDER,
			Mock: func() {
				mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(data.TestGetItemOutputProvider, nil)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			switch tc.Type {
			case CONTRACT:
				//act
				result, err := GetEntity[models.Contract](testDBClient, testKey)
				//assert
				assert.Nil(t, err)
				assert.NotNil(t, result)
			case TARIFF:
				//act
				result, err := GetEntity[models.Tariff](testDBClient, testKey)
				//assert
				assert.Nil(t, err)
				assert.NotNil(t, result)
			case PROVIDER:
				//act
				result, err := GetEntity[models.Provider](testDBClient, testKey)

				//assert
				assert.Nil(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestUnit_GetEntity_Negative(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)
	mockDBManager.EXPECT().GetItem(gomock.Any(), gomock.Any()).AnyTimes().Return(&dynamodb.GetItemOutput{}, errors.New(constants.ResourceNotFound))

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Negative Test Contract Resource Not Found",
			Type: CONTRACT,
		},
		{
			Name: "Negative Test Tariff Resource Not Found",
			Type: TARIFF,
		},
		{
			Name: "Negative Test Provider Resource Not Found",
			Type: PROVIDER,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			switch tc.Type {
			case CONTRACT:
				//act
				result, err := GetEntity[models.Contract](testDBClient, testKey)
				//assert
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), constants.ResourceNotFound)
				assert.Nil(t, result)
			case TARIFF:
				//act
				result, err := GetEntity[models.Tariff](testDBClient, testKey)
				//assert
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), constants.ResourceNotFound)
				assert.Nil(t, result)
			case PROVIDER:
				//act
				result, err := GetEntity[models.Provider](testDBClient, testKey)
				//assert
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), constants.ResourceNotFound)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUnit_PutEntity(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)
	mockDBManager.EXPECT().PutItem(gomock.Any(), gomock.Any()).AnyTimes().Return(&dynamodb.PutItemOutput{}, nil)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Positive Test Contract",
			Type: CONTRACT,
		},
		{
			Name: "Positive Test Tariff",
			Type: TARIFF,
		},
		{
			Name: "Positive Test Provider",
			Type: PROVIDER,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			var err error
			//act
			switch tc.Type {
			case CONTRACT:
				err = PutEntity(testDBClient, models.Contract{})
			case TARIFF:
				err = PutEntity(testDBClient, models.Tariff{})
			case PROVIDER:
				err = PutEntity(testDBClient, models.Provider{})
			}
			//assert
			assert.Nil(t, err)
		})
	}
}

func TestUnit_PutEntity_Negative(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)
	mockDBManager.EXPECT().PutItem(gomock.Any(), gomock.Any()).AnyTimes().Return(&dynamodb.PutItemOutput{}, errors.New(constants.ResourceNotFound))

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Negative Test Contract Resource Not Found",
			Type: CONTRACT,
		},
		{
			Name: "Negative Test Tariff Resource Not Found",
			Type: TARIFF,
		},
		{
			Name: "Negative Test Provider Resource Not Found",
			Type: PROVIDER,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			var err error
			//act
			switch tc.Type {
			case CONTRACT:
				err = PutEntity(testDBClient, models.Contract{})
			case TARIFF:
				err = PutEntity(testDBClient, models.Tariff{})
			case PROVIDER:
				err = PutEntity(testDBClient, models.Provider{})
			}
			//assert
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), constants.ResourceNotFound)
		})
	}
}

func TestUnit_DeleteEntity(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Positive Test",
			Type: POSITIVE,
			Mock: func() {
				mockDBManager.EXPECT().DeleteItem(gomock.Any(), gomock.Any()).Return(&dynamodb.DeleteItemOutput{}, nil)
			},
		},
		{
			Name: "Negative Test",
			Type: NEGATIVE,
			Mock: func() {
				mockDBManager.EXPECT().DeleteItem(gomock.Any(), gomock.Any()).Return(&dynamodb.DeleteItemOutput{}, errors.New(constants.ResourceNotFound))
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			//act
			switch tc.Type {
			case POSITIVE:
				err := DeleteEntity(testDBClient, testKey)
				//assert
				assert.Nil(t, err)
			case NEGATIVE:
				err := DeleteEntity(testDBClient, testKey)
				//assert
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), constants.ResourceNotFound)
			}
		})
	}
}

func TestUnit_UpdateEntity(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	update := expression.Set(expression.Name("TestPath"), expression.Value("TestData"))
	expr, _ := expression.NewBuilder().WithUpdate(update).Build()

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Positive Test",
			Type: POSITIVE,
			Mock: func() {
				mockDBManager.EXPECT().UpdateItem(gomock.Any(), gomock.Any()).Return(&dynamodb.UpdateItemOutput{}, nil)
			},
		},
		{
			Name: "Negative Test Resource Not Found",
			Type: NEGATIVE,
			Mock: func() {
				mockDBManager.EXPECT().UpdateItem(gomock.Any(), gomock.Any()).Return(&dynamodb.UpdateItemOutput{}, errors.New(constants.ResourceNotFound))
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			switch tc.Type {
			case POSITIVE:
				err := UpdateEntity(testDBClient, testKey, expr)

				//assert
				assert.Nil(t, err)
			case NEGATIVE:
				err := UpdateEntity(testDBClient, testKey, expr)

				//assert
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), constants.ResourceNotFound)
			}
		})
	}
}

func Test_QueryEntities(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Positive Test Contract",
			Type: CONTRACT,
			Mock: func() {
				mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(data.TestContractQueryOutput, nil)
			},
		},
		{
			Name: "Positive Test Tariff",
			Type: TARIFF,
			Mock: func() {
				mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(data.TestTariffQueryOutput, nil)
			},
		},
		{
			Name: "Positive Test Provider",
			Type: PROVIDER,
			Mock: func() {
				mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(data.TestProviderQueryOutput, nil)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			switch tc.Type {
			case CONTRACT:
				contracts, err := QueryEntities[models.Contract](testDBClient, data.TestPartitionId, data.TestSortKey)

				//assert
				assert.Nil(t, err)
				assert.NotNil(t, contracts)
			case TARIFF:
				tariffs, err := QueryEntities[models.Tariff](testDBClient, data.TestPartitionId, data.TestSortKey)

				//assert
				assert.Nil(t, err)
				assert.NotNil(t, tariffs)
			case PROVIDER:
				providers, err := QueryEntities[models.Provider](testDBClient, data.TestPartitionId, data.TestSortKey)

				//assert
				assert.Nil(t, err)
				assert.NotNil(t, providers)
			}
		})
	}
}

func Test_QueryEntities_Negative(t *testing.T) {
	//arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockDBManager := dbtesting.NewMockDynamoDBManager(mockController)

	testDBClient := DBClient{
		DynamoDBClient: mockDBManager,
		TableName:      "TestTableName",
		PartitionKey:   "TestPartitionKey",
		SortKey:        "TestSortKey",
	}

	testcases := []TestCase{
		{
			Name: "Negative Test Contract",
			Type: CONTRACT,
			Mock: func() {
				mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{}, errors.New("DatabaseError"))
			},
		},
		{
			Name: "Negative Test Tariff",
			Type: TARIFF,
			Mock: func() {
				mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{}, errors.New("DatabaseError"))
			},
		},
		{
			Name: "Negative Test Provider",
			Type: PROVIDER,
			Mock: func() {
				mockDBManager.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{}, errors.New("DatabaseError"))
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			switch tc.Type {
			case CONTRACT:
				contracts, err := QueryEntities[models.Contract](testDBClient, data.TestPartitionId, data.TestSortKey)

				//assert
				assert.NotNil(t, err)
				assert.Nil(t, contracts)
			case TARIFF:
				tariffs, err := QueryEntities[models.Tariff](testDBClient, data.TestPartitionId, data.TestSortKey)

				//assert
				assert.NotNil(t, err)
				assert.Nil(t, tariffs)
			case PROVIDER:
				providers, err := QueryEntities[models.Provider](testDBClient, data.TestPartitionId, data.TestSortKey)

				//assert
				assert.NotNil(t, err)
				assert.Nil(t, providers)
			}
		})
	}
}
