package service

import (
	"net/http"
)

func (s *service) ListenAndServe() error {

	http.Handle("/", s.authorizer.Authenticate(s.handleProxy()))
	http.Handle("/command/v1/", s.authorizer.Authenticate(s.handleCommand()))

	s.httpServer = &http.Server{Addr: ":80", Handler: http.DefaultServeMux}
	return s.httpServer.ListenAndServe()
}
