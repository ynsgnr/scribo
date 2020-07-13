package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/config"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/controller"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/service"
)

const (
	ServiceName = "file-converter"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	log.SetOutput(logger.New(ses, ServiceName))
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

	controller := controller.NewController(nil)
	s := service.NewService(c, p, controller, cfg.FileTopic)
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