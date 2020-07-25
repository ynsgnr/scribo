package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/gateway/internal/commander"
)

func (s *service) handleCommand() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Printf(logger.Trace, "handleCommand: wrong method: %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte(fmt.Sprintf("wrong method, %s is accepted", http.MethodPost)))
			return
		}
		externalUserID, ok := r.Context().Value(authenticator.HttpUserIDHeader).(string)
		if !ok || externalUserID == "" {
			logger.Printf(logger.Error, "handleCommand: internal user id context is empty")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("internal id auth failed"))
			return
		}
		if r.URL.Path != fmt.Sprintf("/command/v1/user/%s/command", externalUserID) {
			logger.Printf(logger.Trace, "handleCommand: wrong path: %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		internalUserID, ok := r.Context().Value(authenticator.HttpInternalUserIDHeader).(string)
		if !ok || internalUserID == "" {
			logger.Printf(logger.Error, "handleCommand: internal user id context is empty")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("internal id auth failed"))
		}
		eventType := event.Type(r.Header.Get(event.TypeIdentifier))
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Printf(logger.Error, "handleCommand: reading body: %+v", err)
			_, _ = w.Write([]byte("unexpected error"))
			return
		}
		logger.Printf(logger.Trace, "customer %s commanded %s with data: %s", internalUserID, eventType, string(body))
		switch eventType {
		case event.TypeAddDevice:
			s.unmarshalAndSend(w, event.TypeAddDevice, &commander.AddDevice{}, internalUserID, body)
		case event.TypeSend2Device:
			s.unmarshalAndSend(w, event.TypeSend2Device, &commander.SyncDevice{}, internalUserID, body)
		default:
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("unknown event type"))
		}
	})
}

func (s *service) unmarshalAndSend(w http.ResponseWriter, eventType event.Type, template commander.DataInterface, key string, body []byte) {
	err := json.Unmarshal(body, &template)
	if err != nil {
		logger.Printf(logger.Error, "handleCommand %s: UnmarshalAndSend: %s: %+v", eventType, string(body), err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("unexpected error while parsing"))
		return
	}
	protoObject, err := template.ToProto()
	if err != nil {
		logger.Printf(logger.Error, "handleCommand %s: template.ToProto: %s: %+v", eventType, string(body), err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("unexpected error while parsing"))
		return
	}
	err = s.commander.Send(commander.Command{
		EventType: eventType,
		Key:       key,
		Data:      protoObject,
	})
	if err != nil {
		logger.Printf(logger.Error, "handleCommand %s: UnmarshalAndSend: %s: commander.Send: %+v", eventType, string(body), err)
		_, _ = w.Write([]byte("unexpected error"))
		return
	}
}
