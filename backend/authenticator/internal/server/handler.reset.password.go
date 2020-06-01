package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) handleResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s.writeError(NotImplemented{}, w)
}
