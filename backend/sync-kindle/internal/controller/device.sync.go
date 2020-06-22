package controller

import (
	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
)

func (c *controller) SyncDevice(sync2Device *device.Sync2Device) (*mail.SendMail, error) {
	email, err := c.repository.Read(sync2Device.DeviceID)
	if err != nil {
		return nil, errors.Wrapf(err, "controller: SyncDevice: repository.Read: for %s", sync2Device.DeviceID)
	}
	return &mail.SendMail{
		MailID:             sync2Device.SyncID,
		To:                 email,
		AttachmentLocation: sync2Device.FileLocation,
	}, nil
}

func (c *controller) SyncMailSend(sendMail *mail.SendMail) *device.Sync2Device {
	return &device.Sync2Device{
		SyncID: sendMail.MailID,
	}
}
