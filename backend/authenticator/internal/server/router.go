package server

import "net/http"

func (s *server) ListenAndServe() error {
	s.router.POST("/auth/user", s.handleSignUp)
	s.router.PUT("/auth/user", s.handleSignIn)
	s.router.PATCH("/auth/user", s.handleResetPassword)
	s.router.DELETE("/auth/user", s.handleSignOut)
	s.router.GET("/auth/user", s.handleValidate)

	return http.ListenAndServe(":81", s.router)
}
