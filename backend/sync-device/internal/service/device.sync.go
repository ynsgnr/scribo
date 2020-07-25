package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
	"google.golang.org/protobuf/proto"
)

func (s *service) syncDevice(key []byte, value []byte) {
	sync2Device := &device.Sync2Device{}
	err := proto.Unmarshal(value, sync2Device)
	if err != nil {
		logger.Printf(logger.Error, "syncDevice: unmarshal: %s", err.Error())
		return
	}
	fileConvert, deviceSync, err := s.controller.SyncDevice(string(key), sync2Device)
	if err != nil {
		logger.Printf(logger.Error, err.Error())
		return
	}
	if deviceSync != nil {
		msg, err := proto.Marshal(sync2Device)
		if err != nil {
			logger.Printf(logger.Error, "syncDevice: controller.SyncDevice: Marshal: %s for %s", err.Error(), sync2Device.SyncID)
			return
		}
		s.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &s.syncDeviceTopic, Partition: kafka.PartitionAny},
			Key:            key,
			Value:          msg,
			Headers: []kafka.Header{{
				Key:   string(event.TypeIdentifier),
				Value: []byte(event.TypeSend2Device),
			}},
		}, nil)
	} else if fileConvert != nil {
		msg, err := proto.Marshal(fileConvert)
		if err != nil {
			logger.Printf(logger.Error, "syncDevice: controller.SyncDevice: Marshal: %s for %s", err.Error(), fileConvert.FileID)
			return
		}
		s.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &s.fileConvertTopic, Partition: kafka.PartitionAny},
			Key:            key,
			Value:          msg,
			Headers: []kafka.Header{{
				Key:   string(event.TypeIdentifier),
				Value: []byte(event.TypeConvertFile),
			}},
		}, nil)
	}
}

func (s *service) syncDeviceSuccessfull(key []byte, value []byte) {
	sync2Device := &device.Sync2Device{}
	err := proto.Unmarshal(value, sync2Device)
	if err != nil {
		logger.Printf(logger.Error, "syncDevice: unmarshal: %s", err.Error())
		return
	}
	err = s.controller.SyncDeviceSuccessfull(string(key), sync2Device)
	if err != nil {
		logger.Printf(logger.Error, err.Error())
		return
	}
}

func (s *service) convertFileSuccess(key []byte, value []byte) {
	convertFile := &file.ConvertFile{}
	err := proto.Unmarshal(value, convertFile)
	if err != nil {
		logger.Printf(logger.Error, "convertFileSuccess: unmarshal: %s", err.Error())
		return
	}
	sync2Device, err := s.controller.ConvertFileSuccessfull(string(key), convertFile)
	if err != nil {
		logger.Printf(logger.Error, "convertFileSuccess: controller.ConvertFileSuccessfull: %s", err.Error())
		return
	}
	msg, err := proto.Marshal(sync2Device)
	if err != nil {
		logger.Printf(logger.Error, "convertFileSuccess: controller.AddDevice: Marshal: %s for %s", err.Error(), convertFile.FileID)
		return
	}
	s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.syncDeviceTopic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(event.TypeSend2DeviceSuccess),
		}},
	}, nil)
}
