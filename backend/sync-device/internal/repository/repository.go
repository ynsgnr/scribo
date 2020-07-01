package repository

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
)

type Device struct {
	device.AddDevice
	UserID string           `json:"userID"`
	Send   map[string]*Send `json:"send"`
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
	State  State  `json:"state"`
}

type Interface interface {
	WriteDevice(*Device) error
	WriteSend(*Send) error
	ReadDevices(userID string) ([]*Device, error)
	DeleteDevice(device *Device) error
}
