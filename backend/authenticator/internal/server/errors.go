package server

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type UserExistsException struct{ error }

func (UserExistsException) Error() string {
	return "UserExistsException"
}

type AuthChallengeException struct{ error }

func (AuthChallengeException) Error() string {
	return "AuthChallengeException"
}

type JSONMarshallError struct{ error }

func (JSONMarshallError) Error() string {
	return "JSONMarshallError"
}

type NotImplemented struct{ error }

func (NotImplemented) Error() string {
	return "NotImplemented"
}

func (s *server) writeError(err error, w http.ResponseWriter) {
	switch err.(type) {
	case *cognitoidentityprovider.ResourceNotFoundException:
		w.WriteHeader(http.StatusNotFound)
	case *cognitoidentityprovider.InvalidParameterException:
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid paramater"))
	case *cognitoidentityprovider.UserNotFoundException:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("user not found"))
	case *cognitoidentityprovider.TooManyRequestsException:
		w.WriteHeader(http.StatusTooManyRequests)
	case *cognitoidentityprovider.PasswordResetRequiredException:
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte("password reset required"))
	case *cognitoidentityprovider.UserNotConfirmedException:
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte("user not confirmed"))
	case *cognitoidentityprovider.NotAuthorizedException:
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("not authorized"))
	case *cognitoidentityprovider.InvalidPasswordException:
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("password policy did not conform"))
	case UserExistsException:
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte("user already exists"))
	case AuthChallengeException:
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("account is blocked, auth challenge needs to be completed"))
	case JSONMarshallError:
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	case NotImplemented:
		w.WriteHeader(http.StatusNotImplemented)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
