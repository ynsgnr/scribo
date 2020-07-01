package repository

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
)

type Device struct {
	device.AddDevice
	FileTarget string           `json:"fileTarget"`
	UserID     string           `json:"userID"`
	Send       map[string]*Send `json:"send"`
}

type State string

const (
	StateWaitingFileConvert State = "waitingFileConvert"
	StateWaitingSync        State = "waitingSync"
	StateDone               State = "done"
)

type Send struct {
	device.Sync2Device
	UserID string `json:"userID"`
	State  State  `json:"syncState"`
}

type Interface interface {
	WriteDevice(*Device) error
	WriteSend(*Send) error
	GetSendByFileID(string, string) (*Send, error)
	GetDevice(string, string) (*Device, error)
	UpdateSendState(userID string, deviceID string, syncID string, state State) error
	ReadDevices(string) ([]*Device, error)
	DeleteDevice(*Device) error
}
