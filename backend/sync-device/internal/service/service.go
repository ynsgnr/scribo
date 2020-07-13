package service

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/controller"
	"github.com/ynsgnr/scribo/backend/sync-device/internal/repository"
)

type Interface interface {
	Run()
	Shutdown(time.Duration)
}

func NewService(consumer *kafka.Consumer, producer *kafka.Producer, controller controller.Interface, repository repository.Interface, addDeviceTopic, syncDeviceTopic, commandTopic, fileConvertTopic string) Interface {
	return &service{
		consumer:   consumer,
		producer:   producer,
		controller: controller,
		repository: repository,

		commandTopic:     commandTopic,
		fileConvertTopic: fileConvertTopic,
		addDeviceTopic:   addDeviceTopic,
		syncDeviceTopic:  syncDeviceTopic,

		keepRunning: true,
	}
}

type service struct {
	consumer   *kafka.Consumer
	producer   *kafka.Producer
	controller controller.Interface
	repository repository.Interface

	router     *httprouter.Router
	httpServer *http.Server

	commandTopic     string
	fileConvertTopic string
	addDeviceTopic   string
	syncDeviceTopic  string

	keepRunning bool
}

func (s *service) Run() {
	err := s.consumer.SubscribeTopics([]string{s.commandTopic, s.syncDeviceTopic, s.fileConvertTopic}, nil)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if rec := recover(); rec != nil {
				logger.Printf(logger.Error, "%+v", rec)
			}
		}()
		err := s.ListenAndServe()
		if err != nil {
			logger.Printf(logger.Error, err.Error())
		}
	}()
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
	wg.Wait()
}

func (s *service) Shutdown(timeout time.Duration) {
	s.consumer.Close()
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logger.Printf(logger.Error, err.Error())
		}
	}
	timeoutMS := int(timeout.Milliseconds())
	if timeoutMS < 0 {
		//Overflow check
		panic("service: shutdown: given timeout is overflowed for int value")
	}
	s.producer.Flush(timeoutMS)
}
