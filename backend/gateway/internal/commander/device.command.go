package commander

import (
	"fmt"

	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
)

type DataInterface interface {
	ToProto() (interface{}, error)
}

type DeviceType string

const (
	Default        DeviceType = "default"
	Kindle         DeviceType = "kindle"
	KindleEmailKey            = "kindleEmail"
)

var deviceMap = map[DeviceType]int32{
	Default: 0,
	Kindle:  1,
}

func (d DeviceType) ToInt() int32 {
	return deviceMap[d]
}

type AddDevice struct {
	Name string            `json:"name"`
	Type DeviceType        `json:"type"`
	Data map[string]string `json:"data"`
}

func (d *AddDevice) ToProto() (interface{}, error) {
	protoDevice := &device.AddDevice{
		DeviceName: d.Name,
		DeviceType: device.DeviceType(d.Type.ToInt()),
	}
	switch d.Type {
	case Kindle:
		kindleEmail, ok := d.Data[KindleEmailKey]
		if !ok {
			return nil, fmt.Errorf("kindleEmail data is required for device type kindle")
		}
		protoDevice.AddKindle = &device.AddKindle{
			KindleEmail: kindleEmail,
		}
	default:
		return nil, fmt.Errorf("unexpected device type")
	}
	return protoDevice, nil
}

type SyncDevice struct {
	DeviceID     string `json:"deviceID"`
	FileLocation string `json:"fileLocation"`
}

func (s *SyncDevice) ToProto() (interface{}, error) {
	return &device.Sync2Device{
		DeviceID:     s.DeviceID,
		FileLocation: s.FileLocation,
	}, nil
}
