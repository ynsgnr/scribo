package server

import "net/http"

func (s *server) ListenAndServe() error {
	s.router.POST("/auth/user", s.handleSignUp)
	s.router.PATCH("/auth/user", s.handleResetPassword)
	s.router.DELETE("/auth/user/session", s.handleSignOut)
	s.router.PUT("/auth/user/session", s.handleSignIn)
	s.router.GET("/auth/user/session", s.handleValidate)

	return http.ListenAndServe(":81", s.router)
}
