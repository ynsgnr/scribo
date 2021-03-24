package repository

import (
	"encoding/json"
	"errors"
)

type DeviceQueryResult struct {
	DeviceName string                      `json:"deviceName"`
	DeviceID   string                      `json:"deviceID"`
	DeviceType DeviceType                  `json:"deviceType"`
	FileTarget string                      `json:"fileTarget"`
	Send       map[string]*SendQueryResult `json:"send"`
}

type SendQueryResult struct {
	SyncID       string `json:"syncID"`
	DeviceID     string `json:"deviceID"`
	FileLocation string `json:"fileLocation"`
	State        State  `json:"syncState"`
}

type DeviceType int32

const (
	Default       DeviceType = 0
	Kindle        DeviceType = 1
	defaultString string     = "default"
	kindleString  string     = "kindle"
)

var deviceStringMap = map[DeviceType]string{
	Default: defaultString,
	Kindle:  kindleString,
}

var deviceIntMap = map[string]DeviceType{
	defaultString: Default,
	kindleString:  Kindle,
}

func (d DeviceType) String() string {
	return deviceStringMap[d]
}

func (d *DeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DeviceType) UnmarshalJSON(b []byte) error {
	value, ok := deviceIntMap[string(b)]
	if !ok {
		return errors.New("unexpected device type")
	}
	*d = value
	return nil
}
