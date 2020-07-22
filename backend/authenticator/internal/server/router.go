package server

import "net/http"

func (s *server) ListenAndServe() error {
	s.router.POST("/auth/v1/user", s.handleSignUp)
	s.router.PATCH("/auth/v1/user", s.handleResetPassword)
	s.router.DELETE("/auth/v1/user/session", s.handleSignOut)
	s.router.PUT("/auth/v1/user/session", s.handleSignIn)
	s.router.GET("/auth/v1/user/session", s.handleValidate)

	return http.ListenAndServe(":81", s.router)
}
