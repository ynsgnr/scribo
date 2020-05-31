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

func (s *server) handleSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var signInRequest authenticator.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&signInRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(s.cognitoClient),
		UserPoolId: aws.String(s.cognitoUserPool),
	}
	if signInRequest.Token == "" {
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
	_, err = s.cognito.AdminInitiateAuth(input)
	if err != nil {
		s.writeError(err, w)
		logger.Printf(logger.Error, "handleSignIn: cognito.AdminInitiateAuth: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
