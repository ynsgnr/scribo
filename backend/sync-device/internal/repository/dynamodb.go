package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

const (
	//devicesTable   = "sync-device"
	dynamoDeviceID = "deviceID"
	dynamoUserID   = "userID"
	dynamoSend     = "send"
)

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

// ReadDevices Read the whole device with send books data
func (ddb *dynamoRepo) ReadDevices(userID string) ([]*Device, error) {
	result, err := ddb.dbClient.Query(&dynamodb.QueryInput{
		TableName:              aws.String(ddb.tableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :i ", dynamoUserID)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":i": {
				S: aws.String(userID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	values := make([]*Device, 0, len(result.Items))
	for _, item := range result.Items {
		value := Device{}
		err = dynamodbattribute.UnmarshalMap(item, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, &value)
	}
	return values, err
}

// WriteDevice create a new device, overrides old device data including send book data
func (ddb *dynamoRepo) WriteDevice(device *Device) error {
	dbItem, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return err
	}
	if device.DeviceID == "" {
		device.DeviceID = uuid.NewV4().String()
	}
	if _, ok := dbItem[dynamoSend]; ok {
		dbItem[dynamoSend] = &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{},
		}
	}

	_, err = ddb.dbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(ddb.tableName),
		Item:      dbItem,
	})
	return err
}

// DeleteDevice Delete a device
func (ddb *dynamoRepo) DeleteDevice(device *Device) error {
	_, err := ddb.dbClient.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			dynamoUserID: {
				S: aws.String(device.UserID),
			},
			dynamoDeviceID: {
				S: aws.String(device.DeviceID),
			},
		},
	})
	return err
}

// WriteSend Create or update a send book data inside a device, if device doesn't exists it will create it
func (ddb *dynamoRepo) WriteSend(send *Send) error {
	dbItem, err := dynamodbattribute.Marshal(send)
	if err != nil {
		return err
	}
	_, err = ddb.dbClient.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:        aws.String(ddb.tableName),
		UpdateExpression: aws.String(fmt.Sprintf("SET %s.%s = :i", dynamoSend, send.SyncID)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":i": dbItem,
		},
		Key: map[string]*dynamodb.AttributeValue{
			dynamoUserID: {
				S: aws.String(send.UserID),
			},
			dynamoDeviceID: {
				S: aws.String(send.DeviceID),
			},
		},
	})
	return err
}
