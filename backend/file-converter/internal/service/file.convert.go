package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
)

func (s *service) convertFile(key []byte, value []byte) {
	fileConvert := &file.ConvertFile{}
	err := proto.Unmarshal(value, fileConvert)
	if err != nil {
		logger.Printf(logger.Error, "convertFile: unmarshal: %s", err.Error())
		return
	}
	sendMail, err := s.controller.ConvertFile(fileConvert)
	if err != nil {
		logger.Printf(logger.Error, err.Error())
		return
	}
	msg, err := proto.Marshal(sendMail)
	if err != nil {
		logger.Printf(logger.Error, "convertFile: controller.ConvertFile: Marshal: %s for %s", err.Error(), fileConvert.FileID)
		return
	}
	s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.fileTopic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(event.TypeSendMail),
		}},
	}, nil)
}
