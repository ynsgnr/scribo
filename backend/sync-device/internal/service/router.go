package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *service) ListenAndServe() error {
	s.router = httprouter.New()

	s.router.GET("sync-device/v1/user/:userID/devices", s.handleGetDevice)

	s.httpServer = &http.Server{Addr: ":80", Handler: s.router}
	return s.httpServer.ListenAndServe()
}
