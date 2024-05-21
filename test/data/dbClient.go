package data

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TestPartitionId = "8eb474f4-3bf9-483c-8c4d-6193a7217fa3"
	TestContractId  = "8b026b56-db5e-4b2a-8d7d-a8b69660977f"
	TestTariffId    = "eb40ecd9-74c9-403c-9e11-33d3f1a26bfe"
	TestProviderId  = "67aed530-e284-4f1a-9dde-833b8f4968d4"
	TestSortKey     = "contract#"
)

var TestAttributeValuesContract = map[string]types.AttributeValue{
	"Partition_Id": &types.AttributeValueMemberS{Value: TestPartitionId},
	"Sort_Key":     &types.AttributeValueMemberS{Value: TestSortKey},
	"Data": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
		"Id":          &types.AttributeValueMemberS{Value: TestContractId},
		"Name":        &types.AttributeValueMemberS{Value: "TestContract"},
		"Description": &types.AttributeValueMemberS{Value: "TestContract Description"},
		"StartDate":   &types.AttributeValueMemberS{Value: "2020-03-24T12:04:18Z"},
		"EndDate":     &types.AttributeValueMemberS{Value: "2022-03-24T12:04:18Z"},
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
