package service

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/sync-mail/internal/controller"
)

type Interface interface {
	Run()
	Shutdown(time.Duration)
}

func NewService(consumer *kafka.Consumer, producer *kafka.Producer, controller controller.Interface, emailTopic string, approveKindlePeriod time.Duration) Interface {
	return &service{
		consumer:            consumer,
		producer:            producer,
		controller:          controller,
		approveKindlePeriod: approveKindlePeriod,

		emailTopic: emailTopic,

		keepRunning: true,
		done:        make(chan bool),
	}
}

type service struct {
	consumer            *kafka.Consumer
	producer            *kafka.Producer
	controller          controller.Interface
	approveKindlePeriod time.Duration

	emailTopic string

	keepRunning bool
	done        chan bool
}

func (s *service) Run() {
	err := s.consumer.SubscribeTopics([]string{s.emailTopic}, nil)
	if err != nil {
		panic(err)
	}

	go s.RunBackground(s.approveKindlePeriod)

	for s.keepRunning {
		msg, err := s.consumer.ReadMessage(-1)
		if err != nil {
			logger.Printf(logger.Error, "ReadMessage: %s", err.Error())
		} else if msg == nil || msg.TopicPartition.Topic == nil || msg.Headers == nil {
			logger.Printf(logger.Error, "ReadMessage: nil topic or nil headers: %+v", msg)
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
	close(s.done)
}
