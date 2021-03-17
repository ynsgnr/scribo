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
	result, err := s.signIn(signInRequest.Base)
	if err != nil {
		err = s.blocker.CheckBlock(string(signInRequest.Email))
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		s.writeError(err, w)
		return
	}
	if result.ChallengeName != nil {
		s.writeError(err, w)
		return
	}
	if result.AuthenticationResult == nil {
		logger.Printf(logger.Error, "handleSignIn: unexpected result: %+v", result)
		s.writeError(authenticator.NotAuthorized, w)
		return
	}
	accessToken := authenticator.Token("")
	refreshToken := authenticator.Token("")
	idToken := authenticator.Token("")
	if result.AuthenticationResult.AccessToken != nil {
		accessToken = authenticator.Token(*result.AuthenticationResult.AccessToken)
	}
	if result.AuthenticationResult.RefreshToken != nil {
		refreshToken = authenticator.Token(*result.AuthenticationResult.RefreshToken)
	}
	if result.AuthenticationResult.IdToken != nil {
		idToken = authenticator.Token(*result.AuthenticationResult.IdToken)
	}
	signInResponse := authenticator.SignInResponse{
		IDToken:      idToken,
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    *result.AuthenticationResult.ExpiresIn,
	}
	body, err := json.Marshal(signInResponse)
	if err != nil {
		s.writeError(errors.New("auth result failed to marshall"), w)
		logger.Printf(logger.Error, " (%s) handleSignIn: signInResponse failed to marshall result: %+v", signInRequest.Email, result)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func (s *server) signIn(input authenticator.Base) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	cognitoInput := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(s.cognitoClient),
		UserPoolId: aws.String(s.cognitoUserPool),
		AuthFlow:   aws.String(cognitoidentityprovider.AuthFlowTypeRefreshToken),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(string(input.Token)),
		},
	}
	if input.Token == "" {
		cognitoInput.AuthFlow = aws.String(cognitoidentityprovider.AuthFlowTypeAdminUserPasswordAuth)
		cognitoInput.AuthParameters = map[string]*string{
			"USERNAME": aws.String(string(input.Email)),
			"PASSWORD": aws.String(string(input.Password)),
		}
		input.Email = "token-sign-in"
	}
	result, err := s.cognito.AdminInitiateAuth(cognitoInput)
	if err != nil {
		logger.Printf(logger.Error, " (%s) signIn: cognito.AdminInitiateAuth: %s", input.Email, err.Error())
		return nil, err
	}
	return result, nil
}
