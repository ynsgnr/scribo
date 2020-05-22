package server

import (
	"github.com/julienschmidt/httprouter"
)

type Server interface {
	ListenAndServe() error
}

type server struct {
	router *httprouter.Router
}

func NewServer() (Server, error) {
	return &server{
		router: httprouter.New(),
	}, nil
}
