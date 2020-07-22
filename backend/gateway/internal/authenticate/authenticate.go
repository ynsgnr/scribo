package authenticate

import (
	"context"
	"net/http"
	"strings"

	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

func NewAuthorizerMiddleware(authenticator authenticator.Interface) *UserAuthenticator {
	return &UserAuthenticator{
		authenticator: authenticator,
	}
}

type UserAuthenticator struct {
	authenticator authenticator.Interface
}

func (am *UserAuthenticator) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.TrimLeft(r.URL.Path, "/")
		service := strings.Split(urlPath, "/")
		if len(service) < 1 {
			logger.Printf(logger.Warning, "someone tries to access %s: unrecognized service", r.URL)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if service[0] == authenticator.ServiceName {
			//allow authenticator access without authorization
			next.ServeHTTP(w, r)
			return
		}
		userSplitted := strings.Split(urlPath, "user/")
		if len(userSplitted) < 2 {
			logger.Printf(logger.Warning, "someone tries to access %s", r.URL)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tokens := strings.Split(r.Header.Get(authenticator.HttpAuthHeader), authenticator.HttpAuthType)
		if len(tokens) < 2 {
			logger.Printf(logger.Trace, "failed to find auth token")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		token := strings.TrimSpace(tokens[1])
		externalID, userID, err := am.authenticator.ValidateToken(token)
		if err != nil {
			logger.Printf(logger.Trace, "failed to authenticate by authenticator: %+v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		urlUserID := strings.Split(userSplitted[1], "/")[0]
		if urlUserID != externalID {
			logger.Printf(logger.Warning, "user %s tries to access %s", userID, r.URL)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), authenticator.HttpUserIDHeader, externalID)
		ctx = context.WithValue(ctx, authenticator.HttpInternalUserIDHeader, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
