package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/device"
	"github.com/ynsgnr/scribo/backend/gateway/internal/commander"
)

func (s *service) handleCommand() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Printf(logger.Trace, "handleCommand: wrong method: %s", r.Method)
			_, _ = w.Write([]byte(fmt.Sprintf("wrong method, %s is accepted", http.MethodPost)))
			return
		}
		eventType := event.Type(r.Header.Get(event.TypeIdentifier))
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Printf(logger.Error, "handleCommand: reading body: %+v", err)
			_, _ = w.Write([]byte("unexpected error"))
			return
		}
		internalUserID, ok := r.Context().Value(authenticator.HttpInternalUserIDHeader).(string)
		if !ok || internalUserID == "" {
			logger.Printf(logger.Error, "handleCommand: internal user id context is empty")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("internal id auth failed"))
		}
		logger.Printf(logger.Trace, "customer %s commanded %s with data: %s", internalUserID, eventType, string(body))
		switch eventType {
		case event.TypeAddDevice:
			s.unmarshalAndSend(w, event.TypeAddDevice, device.AddDevice{}, internalUserID, body)
		case event.TypeSend2Device:
			s.unmarshalAndSend(w, event.TypeSend2Device, device.Sync2Device{}, internalUserID, body)
		default:
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("unknown event type"))
		}
	})
}

func (s *service) unmarshalAndSend(w http.ResponseWriter, eventType event.Type, template interface{}, key string, body []byte) {
	err := json.Unmarshal(body, &template)
	if err != nil {
		logger.Printf(logger.Error, "handleCommand %s: UnmarshalAndSend: %s", event.TypeAddDevice, string(body))
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unexpected error"))
		return
	}
	err = s.commander.Send(commander.Command{
		EventType: eventType,
		Key:       key,
		Data:      template,
	})
	if err != nil {
		logger.Printf(logger.Error, "handleCommand %s: UnmarshalAndSend: %s: commander.Send:", event.TypeAddDevice, string(body))
		_, _ = w.Write([]byte("unexpected error"))
		return
	}
}
