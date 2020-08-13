package service

import (
	"net/http"
)

func (s *service) ListenAndServe() error {
	// reverse proxy for other services
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", s.crossOriginAllow)
		w.Header().Add("Access-Control-Allow-Credentials", s.crossOriginAllowCredentials)
		w.Header().Add("Access-Control-Allow-Methods", s.crossOriginAllowMethods)
		w.Header().Add("Access-Control-Allow-Headers", s.crossOriginAllowHeaders)
		w.Header().Add("Access-Control-Expose-Headers", s.crossOriginExposeHeaders)
		s.authorizer.Authenticate(s.handleProxy()).ServeHTTP(w, r)
	}))
	// /command/v1/user/{{user-id}}/command
	http.Handle("/command/v1/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", s.crossOriginAllow)
		w.Header().Add("Access-Control-Allow-Credentials", s.crossOriginAllowCredentials)
		w.Header().Add("Access-Control-Allow-Methods", s.crossOriginAllowMethods)
		w.Header().Add("Access-Control-Allow-Headers", s.crossOriginAllowHeaders)
		s.authorizer.Authenticate(s.handleCommand()).ServeHTTP(w, r)
	}))

	s.httpServer = &http.Server{Addr: ":80", Handler: http.DefaultServeMux}
	return s.httpServer.ListenAndServe()
}
