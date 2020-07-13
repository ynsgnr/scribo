package controller

import (
	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/repository"
)

var targetMap = map[device.DeviceType]string{
	device.DeviceType_KINDLE: "mobi",
}

func (c *controller) AddDevice(userID string, addDevice *device.AddDevice) (*device.AddDevice, error) {
	fileTarget := targetMap[addDevice.DeviceType]
	err := c.repository.WriteDevice(&repository.Device{
		AddDevice:  *addDevice,
		FileTarget: fileTarget,
		UserID:     userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "controller: AddDevice: repository.WriteDevice: %w")
	}
	return addDevice, nil
}
