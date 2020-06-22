package repository

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	dynamoKeyAttribute   = "key"
	dynamoValueAttribute = "value"
)

var NotFoundException = errors.New("not found")

func NewDynamoRepo(dbClient *dynamodb.DynamoDB, tableName string) Interface {
	return &dynamoRepo{
		dbClient:  dbClient,
		tableName: tableName,
	}
}

type dynamoRepo struct {
	dbClient  *dynamodb.DynamoDB
	tableName string
}

func (ddb *dynamoRepo) Read(key string) (string, error) {
	result, err := ddb.dbClient.GetItem(&dynamodb.GetItemInput{
		TableName:      aws.String(ddb.tableName),
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			dynamoKeyAttribute: {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		return "", err
	}
	value, ok := result.Item[dynamoValueAttribute]
	if !ok {
		return "", NotFoundException
	}
	return *value.S, nil
}

func (ddb *dynamoRepo) Write(key, value string) error {
	_, err := ddb.dbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(ddb.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			dynamoKeyAttribute: {
				S: aws.String(key),
			},
			dynamoValueAttribute: {
				S: aws.String(value),
			},
		},
	})
	return err
}
