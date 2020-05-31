package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/authenticator/internal/config"
	"github.com/ynsgnr/scribo/backend/authenticator/internal/server"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	log.SetOutput(logger.New(ses, authenticator.ServiceName))
	logger.Print(logger.Info, "starting service")
	s, err := server.NewServer(cognitoidentityprovider.New(ses), cfg)
	if err != nil {
		panic(err)
	}
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
	logger.Print(logger.Info, "stoping service")
}
