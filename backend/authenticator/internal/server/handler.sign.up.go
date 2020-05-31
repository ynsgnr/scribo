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

func (s *server) handleSignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var signUpRequest authenticator.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&signUpRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, err = s.cognito.AdminCreateUser(&cognitoidentityprovider.AdminCreateUserInput{
		MessageAction:     aws.String(cognitoidentityprovider.MessageActionTypeSuppress), //Resend the sign up message if user already exists
		Username:          aws.String(string(signUpRequest.Email)),
		UserPoolId:        aws.String(s.cognitoUserPool),
		TemporaryPassword: aws.String(string(signUpRequest.Password)),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("False"),
			},
		},
	})
	if _, ok := err.(*cognitoidentityprovider.UsernameExistsException); !ok && err != nil {
		s.writeError(err, w)
		logger.Printf(logger.Error, "handleSignUp: cognito.AdminCreateUser: %s", err.Error())
		return
	}
	sesData, err := s.cognito.AdminInitiateAuth(&cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:   aws.String(cognitoidentityprovider.AuthFlowTypeAdminUserPasswordAuth),
		ClientId:   aws.String(s.cognitoClient),
		UserPoolId: aws.String(s.cognitoUserPool),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(string(signUpRequest.Email)),
			"PASSWORD": aws.String(string(signUpRequest.Password)),
		},
	})
	if _, ok := err.(*cognitoidentityprovider.NotAuthorizedException); ok {
		s.writeError(err, w)
		return
	} else if err != nil {
		s.writeError(err, w)
		logger.Printf(logger.Error, "handleSignUp: cognito.AdminInitiateAuth: %s", err.Error())
		return
	}
	_, err = s.cognito.AdminRespondToAuthChallenge(&cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		ChallengeName: aws.String(cognitoidentityprovider.ChallengeNameTypeNewPasswordRequired),
		ClientId:      aws.String(s.cognitoClient),
		UserPoolId:    aws.String(s.cognitoUserPool),
		ChallengeResponses: map[string]*string{
			"USERNAME":     aws.String(string(signUpRequest.Email)),
			"NEW_PASSWORD": aws.String(string(signUpRequest.Password)),
		},
		Session: sesData.Session,
	})
	if _, ok := err.(*cognitoidentityprovider.InvalidParameterException); ok {
		s.writeError(UserExistsException{}, w)
		return
	} else if err != nil {
		s.writeError(err, w)
		logger.Printf(logger.Error, "handleSignUp: cognito.AdminRespondToAuthChallenge: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
