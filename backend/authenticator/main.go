package main

import "github.com/ynsgnr/scribo/backend/authenticator/internal/server"

func main() {
	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
