package service

import (
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
)

func (s *service) topics(topic string, key []byte, value []byte, eventType event.Type) {
	logger.Printf(logger.Trace, "received: type: %s, topic: %s, key: %s", eventType, topic, string(key))
	switch eventType {
	case event.TypeUnknown:
	case event.TypeSendMail:
		s.sendMail(key, value)
	}
}
