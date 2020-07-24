package service

import (
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
)

func (s *service) topics(topic string, key []byte, value []byte, eventType event.Type) {
	logger.Printf(logger.Trace, "received: type: %s, topic: %s, key: %s", eventType, topic, string(key))
	switch {
	case eventType == event.TypeUnknown:
	case eventType == event.TypeAddDevice && topic == s.commandTopic:
		s.addDevice(key, value)
	case eventType == event.TypeSend2Device && topic == s.commandTopic:
		s.syncDevice(key, value)
	case eventType == event.TypeConvertFileSuccess && topic == s.fileConvertTopic:
		s.convertFileSuccess(key, value)
	case eventType == event.TypeSend2DeviceSuccess && topic == s.syncDeviceTopic:
		s.syncDeviceSuccessfull(key, value)
	}
}
