package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/controller"
	"google.golang.org/protobuf/proto"
)

func (s *service) addDevice(key []byte, value []byte) {
	addDevice := &device.AddDevice{}
	err := proto.Unmarshal(value, addDevice)
	if err != nil {
		logger.Printf(logger.Error, "addDevice: unmarshal: %s", err.Error())
		return
	}
	addDevice, err = s.controller.AddDevice(addDevice)
	switch {
	case err == nil:
	case errors.Is(err, controller.MissingKindleData):
		logger.Printf(logger.Warning, err.Error())
		return
	default:
		logger.Printf(logger.Error, err.Error())
		return
	}
	msg, err := proto.Marshal(addDevice)
	if err != nil {
		logger.Printf(logger.Error, "addDevice: controller.AddDevice: Marshal: %s for %s", err.Error(), addDevice.DeviceID)
		return
	}
	s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.addDeviceTopic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(event.TypeAddDeviceSuccess),
		}},
	}, nil)
}
