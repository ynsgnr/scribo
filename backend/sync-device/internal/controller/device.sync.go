package controller

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/repository"
)

func (c *controller) SyncDevice(userID string, sync2Device *device.Sync2Device) (*file.ConvertFile, error) {
	device, err := c.repository.GetDevice(userID, sync2Device.DeviceID)
	if err != nil {
		return nil, err
	}
	fileID := uuid.NewV4().String()
	sync2Device.FileID = fileID
	err = c.repository.WriteSend(&repository.Send{
		Sync2Device: *sync2Device,
		UserID:      userID,
		State:       repository.StateWaitingFileConvert,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "controller: SyncDevice: repository.WriteSend: for %s", sync2Device.DeviceID)
	}
	return &file.ConvertFile{
		FileID:       fileID,
		FileLocation: sync2Device.FileLocation,
		Target:       device.FileTarget,
	}, nil
}

func (c *controller) ConvertFileSuccessfull(userID string, convertFile *file.ConvertFile) (*device.Sync2Device, error) {
	send, err := c.repository.GetSendByFileID(userID, convertFile.FileID)
	if err != nil {
		return nil, err
	}
	send.State = repository.StateWaitingSync
	err = c.repository.WriteSend(send)
	return &send.Sync2Device, err
}

func (c *controller) SyncDeviceSuccessfull(userID string, sync2Device *device.Sync2Device) error {
	return c.repository.UpdateSendState(userID, sync2Device.DeviceID, sync2Device.SyncID, repository.StateDone)
}