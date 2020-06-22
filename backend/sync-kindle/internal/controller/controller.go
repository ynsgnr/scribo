package controller

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/repository"
)

type Interface interface {
	AddDevice(*device.AddDevice) (*device.AddDevice, error)
	SyncDevice(*device.Sync2Device) (*mail.SendMail, error)
	SyncMailSend(*mail.SendMail) *device.Sync2Device
}

func NewController(repository repository.Interface) Interface {
	return &controller{
		repository: repository,
	}
}

type controller struct {
	repository repository.Interface
}
