package server

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
)

type Server interface {
	ListenAndServe() error
}

type server struct {
	router  *httprouter.Router
	cognito *cognitoidentityprovider.CognitoIdentityProvider
}

func NewServer(cognito *cognitoidentityprovider.CognitoIdentityProvider) (Server, error) {
	return &server{
		router:  httprouter.New(),
		cognito: cognito,
	}, nil
}
