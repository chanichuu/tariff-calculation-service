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
			&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
				"Id":         &types.AttributeValueMemberS{Value: TestTariffId},
				"Name":       &types.AttributeValueMemberS{Value: "TestTariff"},
				"Currency":   &types.AttributeValueMemberS{Value: "€"},
				"ValidFrom":  &types.AttributeValueMemberS{Value: "2020-03-24T12:04:18Z"},
				"ValidTo":    &types.AttributeValueMemberS{Value: "2022-03-24T12:04:18Z"},
				"TariffType": &types.AttributeValueMemberN{Value: "1"},
			}},
		}},
	}},
}

var TestGetItemOutputContract = &dynamodb.GetItemOutput{
	Item: TestAttributeValuesContract,
}

var TestAttributeValuesTariff = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":         &types.AttributeValueMemberS{Value: TestTariffId},
		"Name":       &types.AttributeValueMemberS{Value: "TestTariff"},
		"Currency":   &types.AttributeValueMemberS{Value: "€"},
		"ValidFrom":  &types.AttributeValueMemberS{Value: "2020-03-24T12:04:18Z"},
		"ValidTo":    &types.AttributeValueMemberS{Value: "2022-03-24T12:04:18Z"},
		"TariffType": &types.AttributeValueMemberN{Value: "1"},
	}},
}

var TestGetItemOutputTariff = &dynamodb.GetItemOutput{
	Item: TestAttributeValuesTariff,
}

var TestAttributeValuesProvider = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":    &types.AttributeValueMemberS{Value: TestTariffId},
		"Name":  &types.AttributeValueMemberS{Value: "TestTariff"},
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
