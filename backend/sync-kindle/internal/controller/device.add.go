package controller

import (
	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
)

func (c *controller) AddDevice(addDevice *device.AddDevice) (*device.AddDevice, error) {
	err := c.repository.Write(addDevice.DeviceID, addDevice.AddKindle.KindleEmail)
	if err != nil {
		return nil, errors.Wrap(err, "controller: AddDevice: %w")
	}
	return addDevice, nil
}
