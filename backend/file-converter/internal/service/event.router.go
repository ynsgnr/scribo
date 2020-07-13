package service

import (
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
)

func (s *service) topics(topic string, key []byte, value []byte, eventType event.Type) {
	switch eventType {
	case event.TypeUnknown:
	case event.TypeConvertFile:
		s.convertFile(key, value)
	}
}
