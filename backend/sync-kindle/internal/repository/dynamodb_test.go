package repository_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/repository"
)

const (
	tableName   = "sync-kindle"
	testKey     = "testKey"
	notFoundKey = "notFoundKey"
	testValue   = "testValue"
)

func TestDynamoDbIntegration(t *testing.T) {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(ses, aws.NewConfig())
	dynamoRepo := repository.NewDynamoRepo(client, tableName)
	err := dynamoRepo.Write(testKey, testValue)
	if err != nil {
		t.Error(err)
		return
	}
	value, err := dynamoRepo.Read(testKey)
	if err != nil {
		t.Error(err)
		return
	}
	if value != testValue {
		t.Errorf("Expected:%s Actual:%s", testValue, value)
		return
	}
}

func TestDynamoDbNotFound(t *testing.T) {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(ses, aws.NewConfig())
	dynamoRepo := repository.NewDynamoRepo(client, tableName)
	_, err := dynamoRepo.Read(notFoundKey)
	switch err {
	case repository.NotFoundException:
	default:
		t.Errorf("Expected repository.NotFoundException, got %s", err.Error())
		return
	}
}
