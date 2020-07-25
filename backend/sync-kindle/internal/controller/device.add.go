package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
)

var MissingKindleData = errors.New("missing kindle data")

func (c *controller) AddDevice(addDevice *device.AddDevice) (*device.AddDevice, error) {
	if addDevice.AddKindle == nil {
		return nil, fmt.Errorf("controller: AddDevice: %w: %+v", MissingKindleData, addDevice)
	}
	err := c.repository.Write(addDevice.DeviceID, addDevice.AddKindle.KindleEmail)
	if err != nil {
		return nil, errors.Wrap(err, "controller: AddDevice: %w")
	}
	return addDevice, nil
}
