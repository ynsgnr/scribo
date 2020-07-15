package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/config"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/controller"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/service"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/storage"
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

	//Check if converter installed
	_, err = exec.LookPath("ebook-convert")
	if err != nil {
		panic(fmt.Errorf("ebook-convert command not found: %w", err))
	}

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

	storage := storage.NewStorageS3(s3manager.NewDownloader(ses), s3manager.NewUploader(ses), cfg.S3Bucket, cfg.TMPFolder)
	controller := controller.NewController(storage)
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
