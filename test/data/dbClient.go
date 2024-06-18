package data

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var TestAttributeValuesContract = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":          &types.AttributeValueMemberS{Value: TestContractId},
		"Name":        &types.AttributeValueMemberS{Value: TestContractName},
		"Description": &types.AttributeValueMemberS{Value: TestContractDescription},
		"StartDate":   &types.AttributeValueMemberS{Value: TestContractStartDate},
		"EndDate":     &types.AttributeValueMemberS{Value: TestContractEndDate},
		"Provider":    &types.AttributeValueMemberS{Value: TestProviderId},
		"Tariffs": &types.AttributeValueMemberL{Value: []types.AttributeValue{
			&types.AttributeValueMemberS{Value: "7c433cd3-f3b0-463b-82c2-24177dd7bfe8"},
			&types.AttributeValueMemberS{Value: "8c433cd3-f3b0-463b-82c2-24177dd7bfe8"},
			&types.AttributeValueMemberS{Value: "9c433cd3-f3b0-463b-82c2-24177dd7bfe8"},
		}},
	}},
}

var TestGetItemOutputContract = &dynamodb.GetItemOutput{
	Item: TestAttributeValuesContract,
}

var TestPutItemOutputContract = &dynamodb.PutItemOutput{
	Attributes: TestAttributeValuesContract,
}

var TestUpdateItemOutputContract = &dynamodb.UpdateItemOutput{
	Attributes: TestAttributeValuesContract,
}

var TestAttributeValuesContractWithoutTariffs = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":          &types.AttributeValueMemberS{Value: TestContractId},
		"Name":        &types.AttributeValueMemberS{Value: TestContractName},
		"Description": &types.AttributeValueMemberS{Value: TestContractDescription},
		"StartDate":   &types.AttributeValueMemberS{Value: TestContractStartDate},
		"EndDate":     &types.AttributeValueMemberS{Value: TestContractEndDate},
		"Provider":    &types.AttributeValueMemberS{Value: TestProviderId},
		"Tariffs":     &types.AttributeValueMemberL{Value: []types.AttributeValue{}},
	}},
}

var TestAttributeValuesTariff = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":         &types.AttributeValueMemberS{Value: TestTariffId},
		"Name":       &types.AttributeValueMemberS{Value: TestTariffName},
		"Currency":   &types.AttributeValueMemberS{Value: TestCurrency},
		"ValidFrom":  &types.AttributeValueMemberS{Value: TestValidFrom},
		"ValidTo":    &types.AttributeValueMemberS{Value: TestValidTo},
		"TariffType": &types.AttributeValueMemberN{Value: "3"},
		"FixedTariff": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"PricePerUnit": &types.AttributeValueMemberN{Value: "64.5"},
		}},
		"DynamicTariff": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"HourlyTariffs": &types.AttributeValueMemberL{Value: []types.AttributeValue{
				&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
					"StartTime": &types.AttributeValueMemberN{Value: TestValidFrom},
					"ValidDays": &types.AttributeValueMemberL{Value: []types.AttributeValue{
						&types.AttributeValueMemberN{Value: "0"},
						&types.AttributeValueMemberN{Value: "1"},
						&types.AttributeValueMemberN{Value: "2"},
					}},
					"PricePerUnit": &types.AttributeValueMemberN{Value: "54.2"},
				}},
			}},
		}},
	}},
}

var TestGetItemOutputTariff = &dynamodb.GetItemOutput{
	Item: TestAttributeValuesTariff,
}

var TestGetQueryOutputTariff = &dynamodb.QueryOutput{
	Items: []map[string]types.AttributeValue{
		TestAttributeValuesTariff,
	},
}

var TestGetItemOutputContractWithoutTariffs = &dynamodb.GetItemOutput{
	Item: TestAttributeValuesContractWithoutTariffs,
}

var TestPutItemOutputTariff = &dynamodb.PutItemOutput{
	Attributes: TestAttributeValuesTariff,
}

var TestUpdateItemOutputTariff = &dynamodb.UpdateItemOutput{
	Attributes: TestAttributeValuesTariff,
}

var TestAttributeValuesProvider = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":    &types.AttributeValueMemberS{Value: TestProviderId},
		"Name":  &types.AttributeValueMemberS{Value: "TestProvider"},
		"Email": &types.AttributeValueMemberS{Value: "test@provider.com"},
		"Address": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"Street":     &types.AttributeValueMemberS{Value: "TestStreet"},
			"PostalCode": &types.AttributeValueMemberS{Value: "107-6001"},
			"City":       &types.AttributeValueMemberS{Value: "Tokyo"},
			"Country":    &types.AttributeValueMemberS{Value: "Japan"},
		}},
	}},
}

var TestGetItemOutputProvider = &dynamodb.GetItemOutput{
	Item: TestAttributeValuesProvider,
}

var TestContractQueryOutput = &dynamodb.QueryOutput{
	Items: []map[string]types.AttributeValue{
		TestAttributeValuesContract,
	},
}

var TestContractQueryOutputWithoutTariffs = &dynamodb.QueryOutput{
	Items: []map[string]types.AttributeValue{
		TestAttributeValuesContractWithoutTariffs,
	},
}

var TestContractQueryOutputPagination = &dynamodb.QueryOutput{
	Items: []map[string]types.AttributeValue{
		TestAttributeValuesContract,
	},
	LastEvaluatedKey: map[string]types.AttributeValue{
		"partitionKey": &types.AttributeValueMemberS{Value: TestPartitionId},
		"sortKey":      &types.AttributeValueMemberS{Value: TestSortKey},
	},
}

var TestTariffQueryOutput = &dynamodb.QueryOutput{
	Items: []map[string]types.AttributeValue{
		TestAttributeValuesTariff,
	},
}

var TestProviderQueryOutput = &dynamodb.QueryOutput{
	Items: []map[string]types.AttributeValue{
		TestAttributeValuesProvider,
	},
}

var TestPutItemOutputProvider = &dynamodb.PutItemOutput{
	Attributes: TestAttributeValuesProvider,
}

var TestUpdateItemOutputProvider = &dynamodb.UpdateItemOutput{
	Attributes: TestAttributeValuesProvider,
}
