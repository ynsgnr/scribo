package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
	"google.golang.org/protobuf/proto"
)

func (s *service) syncDevice(key []byte, value []byte) {
	sync2Device := &device.Sync2Device{}
	err := proto.Unmarshal(value, sync2Device)
	if err != nil {
		logger.Printf(logger.Error, "syncDevice: unmarshal: %s", err.Error())
		return
	}
	sendMail, err := s.controller.SyncDevice(sync2Device)
	if err != nil {
		logger.Printf(logger.Error, err.Error())
		return
	}
	msg, err := proto.Marshal(sendMail)
	if err != nil {
		logger.Printf(logger.Error, "syncDevice: controller.SyncDevice: Marshal: %s for %s", err.Error(), sendMail.MailID)
		return
	}
	s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.emailTopic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(event.TypeSendMail),
		}},
	}, nil)
}

func (s *service) syncMailSuccess(key []byte, value []byte) {
	email := &mail.SendMail{}
	err := proto.Unmarshal(value, email)
	if err != nil {
		logger.Printf(logger.Error, "syncMailSuccess: unmarshal: %s", err.Error())
		return
	}
	addDevice, err := s.controller.SyncMailSend(email)
	if err != nil {
		logger.Printf(logger.Error, "syncMailSuccess: controller.SyncMailSend: %s", err.Error())
		return
	}
	msg, err := proto.Marshal(addDevice)
	if err != nil {
		logger.Printf(logger.Error, "syncMailSuccess: controller.AddDevice: Marshal: %s for %s", err.Error(), addDevice.DeviceID)
		return
	}
	s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.addDeviceTopic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(event.TypeSend2DeviceSuccess),
		}},
	}, nil)
}
