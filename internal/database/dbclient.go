package database

import (
	"context"
	"os"

	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func GetEntity[T any](dbClient DBClient, key map[string]types.AttributeValue) (*T, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dbClient.TableName),
		Key:       key,
	}

	result, err := dbClient.DynamoDBClient.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil || len(result.Item) == 0 {
		return nil, errors.New("ResourceNotFoundError") // todo create proper error
	}

	dbEntity := DBEntity[T]{}
	err = attributevalue.UnmarshalMap(result.Item, &dbEntity)
	if err != nil {
		return nil, err
	}

	return &dbEntity.Data, nil
}

func PutEntity[T any](dbClient DBClient, entity T) error {
	value, err := attributevalue.MarshalMap(entity)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      value,
		TableName: &dbClient.TableName,
	}

	_, err = dbClient.DynamoDBClient.PutItem(context.TODO(), input)
	return err
}

func UpdateEntity[T any](dbClient DBClient, key map[string]types.AttributeValue, expr expression.Expression) error {
	_, err := dbClient.DynamoDBClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 &dbClient.TableName,
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueNone,
	})
	if err != nil {
		// todo properly log this
		//log.Fatalf("failed to update entity")
		return err
	}

	return nil
}

func DeleteEntity(dbClient DBClient, key map[string]types.AttributeValue) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &dbClient.TableName,
		Key:       key,
	}

	_, err := dbClient.DynamoDBClient.DeleteItem(context.TODO(), input)

	return err
}

func BatchWriteEntities(dbClient DBClient, writeRequests []types.WriteRequest) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			dbClient.TableName: writeRequests,
		},
	}

	_, err := dbClient.DynamoDBClient.BatchWriteItem(context.TODO(), input)
	return err
}

func QueryEntities[T any](dbClient DBClient, paritionKey, sortKey string) ([]DBEntity[T], error) {
	keyEx := expression.Key(dbClient.PartitionKey).Equal(expression.Value(paritionKey)).And(expression.KeyBeginsWith(expression.Key(dbClient.SortKey), sortKey))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}

	dbEntity, err := query[DBEntity[T]](dbClient, expr)
	if err != nil {
		return nil, err
	}

	return dbEntity, nil
}

func query[T any](dbClient DBClient, expr expression.Expression) ([]T, error) {
	var response *dynamodb.QueryOutput
	queryResponse := []T{}
	for response == nil || response.LastEvaluatedKey != nil {
		lastEvaluatedKey := map[string]types.AttributeValue{}
		if response == nil {
			lastEvaluatedKey = nil
		} else {
			lastEvaluatedKey = response.LastEvaluatedKey
		}
		response, err := dbClient.DynamoDBClient.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:                 &dbClient.TableName,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
			ExclusiveStartKey:         lastEvaluatedKey,
		})
		if err != nil {
			return nil, err
		}
		queryResponsePage := []T{}

		err = attributevalue.UnmarshalListOfMaps(response.Items, &queryResponsePage)
		if err != nil {
			return nil, err
		}
		queryResponse = append(queryResponse, queryResponsePage...)
	}

	return queryResponse, nil
}
