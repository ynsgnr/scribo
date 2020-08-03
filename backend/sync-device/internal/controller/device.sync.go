package controller

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/repository"
)

func (c *controller) SyncDevice(userID string, sync2Device *device.Sync2Device) (*file.ConvertFile, *device.Sync2Device, error) {
	device, err := c.repository.GetDevice(userID, sync2Device.DeviceID)
	if err != nil {
		return nil, nil, err
	}
	if device.DeviceID == "" {
		return nil, nil, fmt.Errorf("SyncDevice: unknown device for %s", sync2Device.SyncID)
	}
	fileID := uuid.NewV4().String()
	sync2Device.FileID = fileID
	sync2Device.SyncID = uuid.NewV4().String()
	repoSync2Device := &repository.Send{
		Sync2Device: *sync2Device,
		UserID:      userID,
	}
	ext := strings.TrimLeft(filepath.Ext(sync2Device.FileLocation), ".")
	if ext == "" {
		logger.Printf(logger.Error, "no file extension found for sync: %s", sync2Device.SyncID)
		repoSync2Device.State = repository.StateFailed
		repoSync2Device.Notes = "no file extension found"
		err = c.repository.WriteSend(repoSync2Device)
		return nil, nil, err
	}
	if ext == device.FileTarget {
		repoSync2Device.State = repository.StateWaitingSync
		err = c.repository.WriteSend(repoSync2Device)
		return nil, sync2Device, err
	}
	repoSync2Device.State = repository.StateWaitingFileConvert
	err = c.repository.WriteSend(repoSync2Device)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "controller: SyncDevice: repository.WriteSend: for %s", sync2Device.DeviceID)
	}
	return &file.ConvertFile{
		FileID:       fileID,
		FileLocation: sync2Device.FileLocation,
		Target:       device.FileTarget,
	}, nil, nil
}

func (c *controller) ConvertFileSuccessfull(userID string, convertFile *file.ConvertFile) (*device.Sync2Device, error) {
	send, err := c.repository.GetSendByFileID(userID, convertFile.FileID)
	if err != nil {
		return nil, err
	}
	send.State = repository.StateWaitingSync
	err = c.repository.WriteSend(send)
	send.FileLocation = convertFile.FileLocation
	return &send.Sync2Device, err
}

func (c *controller) SyncDeviceSuccessfull(userID string, sync2Device *device.Sync2Device) error {
	return c.repository.UpdateSendState(userID, sync2Device.DeviceID, sync2Device.SyncID, repository.StateDone)
}
