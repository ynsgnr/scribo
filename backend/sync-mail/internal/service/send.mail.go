package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
	"google.golang.org/protobuf/proto"
)

func (s *service) sendMail(key []byte, value []byte) {
	sendMail := &mail.SendMail{}
	err := proto.Unmarshal(value, sendMail)
	if err != nil {
		logger.Printf(logger.Error, "sendMail: unmarshal: %s", err.Error())
		return
	}
	sendMail, err = s.controller.SendMail(sendMail)
	if err != nil {
		logger.Printf(logger.Error, err.Error())
		return
	}
	msg, err := proto.Marshal(sendMail)
	if err != nil {
		logger.Printf(logger.Error, "sendMail: controller.SendMail: Marshal: %s for %s", err.Error(), sendMail.MailID)
		return
	}
	s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.emailTopic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(event.TypeSendMailSuccess),
		}},
	}, nil)
}
