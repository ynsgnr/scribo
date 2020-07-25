package service

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

func (s *service) handleGetDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")
	if userID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	devices, err := s.repository.ReadDevices(userID)
	if err != nil {
		logger.Printf(logger.Error, "repository.ReadDevices: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(devices)
	if err != nil {
		logger.Printf(logger.Error, "repository.ReadDevices: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
