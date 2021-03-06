package repository

import "github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"

type Device struct {
	device.AddDevice
	FileTarget string           `json:"fileTarget"`
	UserID     string           `json:"userID"`
	Send       map[string]*Send `json:"send"`
	Notes      string           `json:"notes"`
}

type State string

const (
	StateWaitingFileConvert State = "waitingFileConvert"
	StateWaitingSync        State = "waitingSync"
	StateFailed             State = "failed"
	StateDone               State = "done"
)

type Send struct {
	device.Sync2Device
	UserID string `json:"userID"`
	State  State  `json:"syncState"`
	Notes  string `json:"notes"`
}
