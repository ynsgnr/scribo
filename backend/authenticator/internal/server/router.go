package server

import "net/http"

func (s *server) ListenAndServe() error {
	s.router.PUT("/auth/user", s.handleSignUp)
	s.router.POST("/auth/user", s.handleSignUp)
	s.router.DELETE("/auth/user", s.handleSignOut)

	return http.ListenAndServe(":8080", s.router)
}
