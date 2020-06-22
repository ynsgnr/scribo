package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/config"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/controller"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/repository"
	"github.com/ynsgnr/scribo/backend/sync-kindle/internal/service"
	"github.com/ynsgnr/scribo/backend/sync-kindle/synckindle"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	log.SetOutput(logger.New(ses, synckindle.ServiceName))
	logger.Print(logger.Info, "starting service")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaEndpoint,
		"group.id":          cfg.KafkaGroupID,
		"auto.offset.reset": cfg.KafkaOffset,
	})
	if err != nil {
		panic(err)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg.KafkaEndpoint})
	if err != nil {
		panic(err)
	}

	repository := repository.NewDynamoRepo(dynamodb.New(ses, aws.NewConfig()), cfg.DynamoTableName)
	controller := controller.NewController(repository)
	s := service.NewService(c, p, controller, cfg.AddDeviceTopic, cfg.SyncDeviceTopic, cfg.EmailTopic)
	OnShutDown(func() { s.Shutdown(time.Second) })
	s.Run()
}

func OnShutDown(f func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		f()
	}()
}
