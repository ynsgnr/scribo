package service

import (
	"net/http"
)

func (s *service) ListenAndServe() error {
	// reverse proxy for other services
	http.Handle("/", s.authorizer.Authenticate(s.handleProxy()))
	// /command/v1/user/{{user-id}}/command
	http.Handle("/command/v1/", s.authorizer.Authenticate(s.handleCommand()))

	s.httpServer = &http.Server{Addr: ":80", Handler: http.DefaultServeMux}
	return s.httpServer.ListenAndServe()
}
