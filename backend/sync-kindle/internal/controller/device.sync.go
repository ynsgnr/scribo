package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
)

var IDSplitter = "#"

func (c *controller) SyncDevice(sync2Device *device.Sync2Device) (*mail.SendMail, error) {
	email, err := c.repository.Read(sync2Device.DeviceID)
	if err != nil {
		return nil, errors.Wrapf(err, "controller: SyncDevice: repository.Read: for %s", sync2Device.DeviceID)
	}
	return &mail.SendMail{
		MailID:             fmt.Sprintf("%s%s%s", sync2Device.DeviceID, IDSplitter, sync2Device.SyncID),
		To:                 email,
		AttachmentLocation: sync2Device.FileLocation,
	}, nil
}

func (c *controller) SyncMailSend(sendMail *mail.SendMail) (*device.Sync2Device, error) {
	ids := strings.Split(sendMail.MailID, IDSplitter)
	if len(ids) != 2 || ids[0] == "" || ids[1] == "" {
		return nil, fmt.Errorf("SyncMailSend: unexpected id for mail: %s", sendMail.MailID)
	}
	return &device.Sync2Device{
		DeviceID: ids[0],
		SyncID:   ids[1],
	}, nil
}
