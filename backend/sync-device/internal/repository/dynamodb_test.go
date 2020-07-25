package repository_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/repository"
)

const (
	tableName = "sync-device"
)

var (
	testDevice = &repository.Device{
		AddDevice: device.AddDevice{
			DeviceName: "testDevice",
			DeviceID:   "testDevice",
			DeviceType: device.DeviceType_KINDLE,
		},
		UserID: "testUser",
		Send:   map[string]*repository.Send{},
	}
	testDevice2 = &repository.Device{
		AddDevice: device.AddDevice{
			DeviceName: "testDevice2",
			DeviceID:   "testDevice2",
			DeviceType: device.DeviceType_KINDLE,
		},
		UserID: "testUser",
		Send:   map[string]*repository.Send{},
	}
	testSend = &repository.Send{
		Sync2Device: device.Sync2Device{
			SyncID:   "testSync",
			DeviceID: testDevice.DeviceID,
			FileID:   "testFile",
		},
		UserID: "testUser",
		State:  repository.StateWaitingSync,
		Notes:  "testNote",
	}
)

func TestDynamoDbCreateDevice(t *testing.T) {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(ses, aws.NewConfig())
	dynamoRepo := repository.NewDynamoRepo(client, tableName)
	err := dynamoRepo.WriteDevice(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	err = dynamoRepo.WriteDevice(testDevice2)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := dynamoRepo.ReadDevices(testDevice.UserID)
	if err != nil {
		t.Error(err)
		return
	}
	if len(resp) < 2 {
		t.Errorf("Unexpected length of devices")
		return
	}
	a, err := json.Marshal(resp)
	if err != nil {
		t.Error(err)
		return
	}
	actual := string(a)
	e, err := json.Marshal(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	expected := string(e)
	if !strings.Contains(actual, expected) {
		t.Errorf("\nExpected:%+v\nIn Actual:%+v", expected, actual)
	}
	e, err = json.Marshal(testDevice2)
	if err != nil {
		t.Error(err)
		return
	}
	expected = string(e)
	if !strings.Contains(actual, expected) {
		t.Errorf("\nExpected:%+v\nIn Actual:%+v", expected, actual)
	}
	err = dynamoRepo.DeleteDevice(testDevice)
	if err != nil {
		t.Error(err)
	}
	err = dynamoRepo.DeleteDevice(testDevice2)
	if err != nil {
		t.Error(err)
	}
}

func TestDynamoDbCreateSend(t *testing.T) {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(ses, aws.NewConfig())
	dynamoRepo := repository.NewDynamoRepo(client, tableName)
	err := dynamoRepo.WriteDevice(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	err = dynamoRepo.WriteSend(testSend)
	if err != nil {
		t.Error(err)
		return
	}
	testDevice.Send = map[string]*repository.Send{
		testSend.SyncID: testSend,
	}
	e, err := json.Marshal(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	expected := string(e)
	resp, err := dynamoRepo.ReadDevices(testDevice2.UserID)
	if err != nil {
		t.Error(err)
		return
	}
	if len(resp) < 1 {
		t.Errorf("Unexpected length of devices")
		return
	}
	a, err := json.Marshal(resp)
	if err != nil {
		t.Error(err)
		return
	}
	actual := string(a)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected:\n%+v\nIn Actual:\n%+v", expected, actual)
	}
	// Update send state
	respSend, err := dynamoRepo.GetSendByFileID(testSend.UserID, testSend.FileID)
	if err != nil {
		t.Error(err)
		return
	}
	e, err = json.Marshal(testSend)
	if err != nil {
		t.Error(err)
		return
	}
	expected = string(e)
	a, err = json.Marshal(respSend)
	if err != nil {
		t.Error(err)
		return
	}
	actual = string(a)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected:\n%+v\nIn Actual:\n%+v", expected, actual)
	}
	respSend.State = repository.StateDone
	err = dynamoRepo.WriteSend(respSend)
	if err != nil {
		t.Error(err)
		return
	}
	testSend.State = repository.StateDone
	testDevice.Send = map[string]*repository.Send{
		testSend.SyncID: testSend,
	}
	e, err = json.Marshal(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	expected = string(e)
	resp, err = dynamoRepo.ReadDevices(testDevice2.UserID)
	if err != nil {
		t.Error(err)
		return
	}
	if len(resp) < 1 {
		t.Errorf("Unexpected length of devices")
		return
	}
	a, err = json.Marshal(resp)
	if err != nil {
		t.Error(err)
		return
	}
	actual = string(a)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected:\n%+v\nIn Actual:\n%+v", expected, actual)
	}

	err = dynamoRepo.DeleteDevice(testDevice)
	if err != nil {
		t.Error(err)
	}
}

func TestDynamoDbGetCreate(t *testing.T) {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(ses, aws.NewConfig())
	dynamoRepo := repository.NewDynamoRepo(client, tableName)
	err := dynamoRepo.WriteDevice(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	err = dynamoRepo.WriteSend(testSend)
	if err != nil {
		t.Error(err)
		return
	}
	testDevice.Send = nil
	e, err := json.Marshal(testDevice)
	if err != nil {
		t.Error(err)
		return
	}
	expected := string(e)
	resp, err := dynamoRepo.GetDevice(testDevice.UserID, testDevice.DeviceID)
	if err != nil {
		t.Error(err)
		return
	}
	a, err := json.Marshal(resp)
	if err != nil {
		t.Error(err)
		return
	}
	actual := string(a)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected:\n%+v\nIn Actual:\n%+v", expected, actual)
		return
	}
	// Update send state
	err = dynamoRepo.UpdateSendState(testSend.UserID, testSend.DeviceID, testSend.SyncID, repository.StateWaitingSync)
	if err != nil {
		t.Error(err)
		return
	}
	testSend.State = repository.StateWaitingSync
	respSend, err := dynamoRepo.GetSendByFileID(testSend.UserID, testSend.FileID)
	if err != nil {
		t.Error(err)
		return
	}
	e, err = json.Marshal(testSend)
	if err != nil {
		t.Error(err)
		return
	}
	expected = string(e)
	a, err = json.Marshal(respSend)
	if err != nil {
		t.Error(err)
		return
	}
	actual = string(a)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected:\n%+v\nIn Actual:\n%+v", expected, actual)
	}
	err = dynamoRepo.DeleteDevice(testDevice)
	if err != nil {
		t.Error(err)
	}
}
