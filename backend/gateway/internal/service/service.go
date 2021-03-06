package service

import (
	"context"
	"net/http"
	"time"

	"github.com/ynsgnr/scribo/backend/common/blocker"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/gateway/internal/authenticate"
	"github.com/ynsgnr/scribo/backend/gateway/internal/commander"
)

type Interface interface {
	Run()
	Shutdown(time.Duration)
}

func NewService(commander commander.Interface, authorizer *authenticate.UserAuthenticator, blocker blocker.Blocker,
	crossOriginAllow string,
	crossOriginAllowCredentials string,
	crossOriginAllowMethods string,
	crossOriginAllowHeaders string,
	crossOriginExposeHeaders string) Interface {
	return &service{
		commander:                   commander,
		authorizer:                  authorizer,
		filter:                      FilterTransport{},
		blocker:                     blocker,
		crossOriginAllow:            crossOriginAllow,
		crossOriginAllowCredentials: crossOriginAllowCredentials,
		crossOriginAllowHeaders:     crossOriginAllowHeaders,
		crossOriginAllowMethods:     crossOriginAllowMethods,
		crossOriginExposeHeaders:    crossOriginExposeHeaders,
	}
}

type service struct {
	commander  commander.Interface
	authorizer *authenticate.UserAuthenticator
	filter     http.RoundTripper
	blocker    blocker.Blocker

	crossOriginAllow            string
	crossOriginAllowCredentials string
	crossOriginAllowMethods     string
	crossOriginAllowHeaders     string
	crossOriginExposeHeaders    string

	httpServer *http.Server
}

func (s *service) Run() {
	defer func() {
		if rec := recover(); rec != nil {
			logger.Printf(logger.Error, "%+v", rec)
		}
	}()
	err := s.ListenAndServe()
	if err != nil {
		logger.Printf(logger.Error, err.Error())
	}
}

func (s *service) Shutdown(timeout time.Duration) {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logger.Printf(logger.Error, err.Error())
		}
	}
}
