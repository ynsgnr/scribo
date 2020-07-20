package server

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
)

func (s *server) handleSignOut(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var signOutRequest authenticator.SignOutRequest
	err := json.NewDecoder(r.Body).Decode(&signOutRequest)
	if err != nil {
		s.writeError(JSONMarshallError{err}, w)
		return
	}
	_, err = s.signIn(&cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(s.cognitoClient),
		UserPoolId: aws.String(s.cognitoUserPool),
		AuthFlow:   aws.String(cognitoidentityprovider.AuthFlowTypeRefreshToken),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(string(signOutRequest.Token)),
		},
	}, signOutRequest.Base)
	if err != nil {
		s.writeError(err, w)
		return
	}
	_, err = s.cognito.AdminUserGlobalSignOut(&cognitoidentityprovider.AdminUserGlobalSignOutInput{
		UserPoolId: aws.String(s.cognitoUserPool),
		Username:   aws.String(string(signOutRequest.Email)),
	})
	if err != nil {
		s.writeError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
