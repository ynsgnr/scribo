package server

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
)

func (s *server) handleResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var resetPassRequest authenticator.ResetPassRequest
	err := json.NewDecoder(r.Body).Decode(&resetPassRequest)
	if err != nil {
		s.writeError(JSONMarshallError{err}, w)
		return
	}
	if resetPassRequest.Code == "" {
		_, err = s.cognito.ForgotPassword(&cognitoidentityprovider.ForgotPasswordInput{
			ClientId: aws.String(s.cognitoClient),
			Username: aws.String(string(resetPassRequest.Email)),
		})
		if err != nil {
			s.blocker.CheckBlock(string(resetPassRequest.Email))
			s.writeError(err, w)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
	_, err = s.cognito.ConfirmForgotPassword(&cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(s.cognitoClient),
		ConfirmationCode: aws.String(string(resetPassRequest.Code)),
		Password:         aws.String(string(resetPassRequest.NewPassword)),
		Username:         aws.String(string(resetPassRequest.Email)),
	})
	if err != nil {
		s.writeError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
