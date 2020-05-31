package server

import "net/http"

func (s *server) ListenAndServe() error {
	s.router.POST("/auth/user", s.handleSignUp)
	s.router.PUT("/auth/user", s.handleSignIn)
	s.router.PATCH("/auth/user", s.handleResetPassword)
	s.router.DELETE("/auth/user", s.handleSignOut)

	return http.ListenAndServe(":80", s.router)
}
