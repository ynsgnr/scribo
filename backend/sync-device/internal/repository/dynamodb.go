package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

const (
	//devicesTable   = "sync-device"
	dynamoItemID   = "itemID"
	dynamoDeviceID = "deviceID"
	dynamoFileID   = "fileID"
	dynamoState    = "syncState"
	dynamoUserID   = "userID"
	dynamoSend     = "send"
	dynamoSendType = "SEND"
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
	values := map[string]*Device{}
	for _, item := range result.Items {
		if _, ok := item[dynamoItemID]; !ok {
			return nil, errors.New("ReadDevices: item is missing deviceID")
		}
		if strings.Contains(*item[dynamoItemID].S, dynamoSendType) {
			value := Send{}
			err = dynamodbattribute.UnmarshalMap(item, &value)
			if err != nil {
				return nil, err
			}
			if _, ok := values[value.DeviceID]; !ok {
				values[value.DeviceID] = &Device{
					Send: map[string]*Send{},
				}
			}
			if values[value.DeviceID].Send == nil {
				values[value.DeviceID].Send = map[string]*Send{}
			}
			values[value.DeviceID].Send[value.SyncID] = &value
		} else {
			value := Device{}
			err = dynamodbattribute.UnmarshalMap(item, &value)
			if err != nil {
				return nil, err
			}
			if _, ok := values[value.DeviceID]; ok {
				value.Send = values[value.DeviceID].Send
			}
			values[value.DeviceID] = &value
		}
	}
	returnValue := make([]*Device, 0, len(values))
	for _, d := range values {
		if d.Send == nil {
			d.Send = map[string]*Send{}
		}
		returnValue = append(returnValue, d)
	}
	return returnValue, err
}

func (ddb *dynamoRepo) GetDevice(userID string, deviceID string) (*Device, error) {
	resp, err := ddb.dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			dynamoUserID: {S: aws.String(userID)},
			dynamoItemID: {S: aws.String(deviceID)},
		},
	})
	if err != nil {
		return nil, err
	}
	value := &Device{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, value)
	return value, err
}

func (ddb *dynamoRepo) UpdateSendState(userID string, deviceID string, syncID string, state State) error {
	value := dynamodb.AttributeValue{S: aws.String(string(state))}
	_, err := ddb.dbClient.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			dynamoUserID: {S: aws.String(userID)},
			dynamoItemID: {S: aws.String(fmt.Sprintf("%s#%s#%s", dynamoSendType, deviceID, syncID))},
		},
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			dynamoState: {
				Value: &value,
			},
		},
	})
	return err
}

// WriteDevice create a new device, overrides old device data including send book data
func (ddb *dynamoRepo) WriteDevice(device *Device) error {
	if device.DeviceID == "" {
		device.DeviceID = uuid.NewV4().String()
	}
	dbItem, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return err
	}
	dbItem[dynamoItemID] = &dynamodb.AttributeValue{
		S: aws.String(device.DeviceID),
	}
	delete(dbItem, dynamoSend)
	_, err = ddb.dbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(ddb.tableName),
		Item:      dbItem,
	})
	return err
}

// DeleteDevice Delete a device
func (ddb *dynamoRepo) DeleteDevice(device *Device) error {
	response, err := ddb.dbClient.Query(&dynamodb.QueryInput{
		TableName:              aws.String(ddb.tableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :userid", dynamoUserID)),
		FilterExpression:       aws.String(fmt.Sprintf("%s = :deviceid", dynamoDeviceID)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userid":   {S: aws.String(device.UserID)},
			":deviceid": {S: aws.String(device.DeviceID)},
		},
	})
	if err != nil {
		return err
	}
	for _, r := range response.Items {
		if _, ok := r[dynamoUserID]; !ok {
			return fmt.Errorf("DeleteDevice: dynamoUserID not found inside item: %+v", r)
		}
		if _, ok := r[dynamoItemID]; !ok {
			return fmt.Errorf("DeleteDevice: dynamoItemID not found inside item: %+v", r)
		}
		_, err = ddb.dbClient.DeleteItem(&dynamodb.DeleteItemInput{
			TableName: aws.String(ddb.tableName),
			Key: map[string]*dynamodb.AttributeValue{
				dynamoUserID: r[dynamoUserID],
				dynamoItemID: r[dynamoItemID],
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteSend Create or update a send book data
func (ddb *dynamoRepo) WriteSend(send *Send) error {
	dbItem, err := dynamodbattribute.MarshalMap(send)
	if err != nil {
		return err
	}
	dbItem[dynamoItemID] = &dynamodb.AttributeValue{
		S: aws.String(fmt.Sprintf("%s#%s#%s", dynamoSendType, send.DeviceID, send.SyncID)),
	}
	_, err = ddb.dbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(ddb.tableName),
		Item:      dbItem,
	})
	return err
}

// Updates state by given fileID
func (ddb *dynamoRepo) GetSendByFileID(userID string, fileID string) (*Send, error) {
	response, err := ddb.dbClient.Query(&dynamodb.QueryInput{
		TableName:              aws.String(ddb.tableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :userid", dynamoUserID)),
		FilterExpression:       aws.String(fmt.Sprintf("%s = :fileid", dynamoFileID)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userid": {S: aws.String(userID)},
			":fileid": {S: aws.String(fileID)},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(response.Items) != 1 {
		return nil, fmt.Errorf("UpdateStateByFileID: unexpected number of items inside query response: %+v", response.Items)
	}
	value := &Send{}
	err = dynamodbattribute.UnmarshalMap(response.Items[0], &value)
	return value, err
}
