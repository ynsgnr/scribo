package server

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
)

func (s *server) handleValidate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := r.Header.Get(authenticator.HttpAuthHeader)
	authData := strings.Split(token, authenticator.HttpAuthType)
	//Validate header
	if len(authData) != 2 || authData[0] != "" || authData[1][0] != ' ' {
		s.writeError(&cognitoidentityprovider.NotAuthorizedException{}, w)
		return
	}
	accessToken := authData[1][1:]
	_, err := s.cognito.GetUser(&cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	})
	if err != nil {
		s.writeError(&cognitoidentityprovider.NotAuthorizedException{}, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
