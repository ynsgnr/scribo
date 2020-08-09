package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common"
)

func (s *server) handleValidate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := r.Header.Get(authenticator.HttpAuthHeader)
	authData := strings.Split(token, authenticator.HttpAuthType)
	//Validate header
	if len(authData) != 2 || authData[0] != "" || authData[1][0] != ' ' {
		s.blocker.CheckBlock(token)
		s.writeError(&cognitoidentityprovider.NotAuthorizedException{}, w)
		return
	}
	accessToken := authData[1][1:]
	response, err := s.cognito.GetUser(&cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	})
	if err != nil || response.Username == nil {
		s.blocker.CheckBlock(token)
		s.writeError(&cognitoidentityprovider.NotAuthorizedException{}, w)
		return
	}
	internalID, externalID, err := common.CalculateIDs(*response.Username, s.internalGeneratorSecret, s.extrenalGeneratorSecret)
	if err != nil {
		s.writeError(errors.New("can not generate id"), w)
		return
	}
	w.Header().Add(authenticator.HttpInternalUserIDHeader, internalID)
	w.Header().Add(authenticator.HttpUserIDHeader, externalID)
	w.WriteHeader(http.StatusNoContent)
}
