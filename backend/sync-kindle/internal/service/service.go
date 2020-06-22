package service

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/controller"
)

type Interface interface {
	Run()
	Shutdown(time.Duration)
}

func NewService(consumer *kafka.Consumer, producer *kafka.Producer, controller controller.Interface, addDeviceTopic, syncDeviceTopic, emailTopic string) Interface {
	return &service{
		consumer:   consumer,
		producer:   producer,
		controller: controller,

		addDeviceTopic:  addDeviceTopic,
		syncDeviceTopic: syncDeviceTopic,
		emailTopic:      emailTopic,

		keepRunning: true,
	}
}

type service struct {
	consumer   *kafka.Consumer
	producer   *kafka.Producer
	controller controller.Interface

	addDeviceTopic  string
	syncDeviceTopic string
	emailTopic      string

	keepRunning bool
}

func (s *service) Run() {
	err := s.consumer.SubscribeTopics([]string{s.addDeviceTopic, s.syncDeviceTopic, s.emailTopic}, nil)
	if err != nil {
		panic(err)
	}

	for s.keepRunning {
		msg, err := s.consumer.ReadMessage(-1)
		if err == nil {
			logger.Printf(logger.Error, "ReadMessage: %s", err.Error())
		} else {
			s.topics(*msg.TopicPartition.Topic, msg.Key, msg.Value, event.GetEventTypeFromHeaders(msg.Headers))
		}
	}
}

func (s *service) Shutdown(timeout time.Duration) {
	s.consumer.Close()
	timeoutMS := int(timeout.Milliseconds())
	if timeoutMS < 0 {
		//Overflow check
		panic("service: shutdown: given timeout is overflowed for int value")
	}
	s.producer.Flush(timeoutMS)
}
