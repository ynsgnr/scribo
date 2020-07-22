package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/gateway/internal/authenticate"
	"github.com/ynsgnr/scribo/backend/gateway/internal/commander"
	"github.com/ynsgnr/scribo/backend/gateway/internal/config"
	"github.com/ynsgnr/scribo/backend/gateway/internal/service"
)

const (
	ServiceName = "gateway"
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

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg.KafkaEndpoint})
	if err != nil {
		panic(err)
	}

	commander := commander.NewKafkaCommander(p, cfg.CommandTopic)
	authenticator := authenticator.NewAuthenticatorSDK("http://authenticator", http.DefaultClient)
	authorizer := authenticate.NewAuthorizerMiddleware(authenticator)
	s := service.NewService(commander, authorizer)
	OnShutDown(func() { s.Shutdown(time.Second) })
	s.Run()
	//Wait for error logs to send
	time.Sleep(time.Second)
}

func OnShutDown(f func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		f()
	}()
}
