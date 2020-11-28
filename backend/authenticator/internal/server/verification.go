package server

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

func (s *server) handleVerification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var verificationRequest authenticator.VerificationRequest
	err := json.NewDecoder(r.Body).Decode(&verificationRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	accessToken := s.getAuthToken(r)
	_, err = s.cognito.VerifyUserAttribute(&cognitoidentityprovider.VerifyUserAttributeInput{
		AccessToken:   aws.String(string(accessToken)),
		AttributeName: aws.String("email"),
		Code:          aws.String(string(verificationRequest.Code)),
	})
	if err != nil {
		s.writeError(err, w)
		logger.Printf(logger.Error, "handleVerification: s.confirmVerificationCode: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *server) handleSendVerificationEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	accessToken := s.getAuthToken(r)
	_, err := s.cognito.GetUserAttributeVerificationCode(&cognitoidentityprovider.GetUserAttributeVerificationCodeInput{
		AccessToken:   aws.String(string(accessToken)),
		AttributeName: aws.String("email"),
	})
	if err != nil {
		s.writeError(err, w)
		logger.Printf(logger.Error, "handleVerification: s.confirmVerificationCode: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
