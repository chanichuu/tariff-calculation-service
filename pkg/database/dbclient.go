package database

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// todo PutItem, Query, etc... calls to DynamoDB that can be reused be repositories
type DBClient struct {
	DynamoDBClient *dynamodb.Client
	TableName      string
	PartitionKey   string
	SortKey        string
}

func NewDBClient() DBClient {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return DBClient{}
	}

	dbClient := dynamodb.NewFromConfig(cfg)
	return DBClient{
		DynamoDBClient: dbClient,
		TableName:      os.Getenv("DYNAMODB_TABLE_NAME"),
		PartitionKey:   os.Getenv("PARTITION_KEY"),
		SortKey:        os.Getenv("SORT_KEY"),
	}
}
