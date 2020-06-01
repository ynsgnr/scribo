package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

func (s *server) handleSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var signInRequest authenticator.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&signInRequest)
	if err != nil {
		s.writeError(JSONMarshallError{err}, w)
		return
	}
	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(s.cognitoClient),
		UserPoolId: aws.String(s.cognitoUserPool),
	}
	if signInRequest.Password != "" {
		input.AuthFlow = aws.String(cognitoidentityprovider.AuthFlowTypeAdminUserPasswordAuth)
		input.AuthParameters = map[string]*string{
			"USERNAME": aws.String(string(signInRequest.Email)),
			"PASSWORD": aws.String(string(signInRequest.Password)),
		}
	} else {
		input.AuthFlow = aws.String(cognitoidentityprovider.AuthFlowTypeRefreshToken)
		input.AuthParameters = map[string]*string{
			"REFRESH_TOKEN": aws.String(string(signInRequest.Token)),
		}
	}
	result, err := s.signIn(input, signInRequest.Base)
	if err != nil {
		s.writeError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	body, err := json.Marshal(result.AuthenticationResult)
	if err != nil {
		s.writeError(errors.New("auth result failed to marshall"), w)
		logger.Printf(logger.Error, " (%s) handleSignIn: result.AuthenticationResult failed to marshall result: %+v", signInRequest.Email, result)
		return
	}
	_, _ = w.Write(body)
}

func (s *server) signIn(input *cognitoidentityprovider.AdminInitiateAuthInput, signInRequest authenticator.Base) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	result, err := s.cognito.AdminInitiateAuth(input)
	if err != nil {
		logger.Printf(logger.Error, " (%s) signIn: cognito.AdminInitiateAuth: %s", signInRequest.Email, err.Error())
		return nil, err
	}
	if result.ChallengeName != nil {
		return nil, AuthChallengeException{}
	}
	if result.AuthenticationResult == nil {
		logger.Printf(logger.Error, " (%s) signIn: result.AuthenticationResult is nul. result: %+v", signInRequest.Email, result)
		return nil, errors.New("auth result is nul")
	}
	return result, nil
}
