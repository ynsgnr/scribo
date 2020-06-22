package service

import (
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
)

func (s *service) topics(topic string, key []byte, value []byte, eventType event.Type) {
	switch eventType {
	case event.TypeUnknown:
	case event.TypeAddDevice:
		s.addDevice(key, value)
	case event.TypeSend2Device:
		s.syncDevice(key, value)
	case event.TypeSendMailSuccess:
		s.syncMailSuccess(key, value)
	}
}
