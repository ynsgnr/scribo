package server

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/internal/config"
	"github.com/ynsgnr/scribo/backend/common/blocker"
)

type Server interface {
	ListenAndServe() error
}

type server struct {
	router                  *httprouter.Router
	cognito                 *cognitoidentityprovider.CognitoIdentityProvider
	cognitoClient           string
	cognitoUserPool         string
	internalGeneratorSecret string
	extrenalGeneratorSecret string

	blocker blocker.Blocker
}

func NewServer(cognito *cognitoidentityprovider.CognitoIdentityProvider, cfg config.Config) (Server, error) {
	return &server{
		router:          httprouter.New(),
		cognito:         cognito,
		cognitoClient:   cfg.ClientId,
		cognitoUserPool: cfg.UserPoolId,
		blocker:         blocker.NewBlocker(cfg.BlockerPeriod, cfg.BlockerCleanupPeriod, cfg.BlockerMaxWait, int64(cfg.BlockerThrottleAfterTries)),
	}, nil
}
