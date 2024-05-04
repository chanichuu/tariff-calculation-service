package database

type DBEntity[T any] struct {
	PartitionKey string `dynamodbav:"Partition_Id"`
	SortKey      string `dynamodbav:"Sort_Key"`
	Data         T      `dynamodbav:"Data"`
}
