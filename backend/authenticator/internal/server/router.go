package server

import "net/http"

func (s *server) ListenAndServe() error {
	s.router.POST("/authenticator/v1/user", s.handleSignUp)
	s.router.PATCH("/authenticator/v1/user", s.handleResetPassword)
	s.router.DELETE("/authenticator/v1/user/session", s.handleSignOut)
	s.router.PUT("/authenticator/v1/user/session", s.handleSignIn)
	s.router.GET("/authenticator/v1/user/session", s.handleValidate)

	return http.ListenAndServe(":80", s.router)
}
