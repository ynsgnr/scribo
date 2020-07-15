package controller

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/repository"
)

type Interface interface {
	AddDevice(string, *device.AddDevice) (*device.AddDevice, error)
	SyncDevice(string, *device.Sync2Device) (*file.ConvertFile, *device.Sync2Device, error)
	ConvertFileSuccessfull(string, *file.ConvertFile) (*device.Sync2Device, error)
	SyncDeviceSuccessfull(string, *device.Sync2Device) error
}

func NewController(repository repository.Interface) Interface {
	return &controller{
		repository: repository,
	}
}

type controller struct {
	repository repository.Interface
}
