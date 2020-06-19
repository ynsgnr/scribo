package event

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Type string

const (
	EventTypeIdentifier = "EventType"

	EventTypeUnknown            Type = "Unknown"
	EventTypeAddDevice          Type = "AddDevice"
	EventTypeAddDeviceSuccess   Type = "AddDeviceSuccess"
	EventTypeSync2Device        Type = "Sync2Device"
	EventTypeSync2DeviceSuccess Type = "Sync2DeviceSuccess"
	EventTypeConvertFile        Type = "ConvertFile"
	EventTypeConvertFileSuccess Type = "ConvertFileSuccess"
	EventTypeSend2Device        Type = "Send2Device"
	EventTypeSend2DeviceSuccess Type = "Send2DeviceSuccess"
	EventTypeSendMail           Type = "SendMail"
	EventTypeSendMailSuccess    Type = "SendMailSuccess"
)

func GetEventTypeFromHeaders(headers []kafka.Header) Type {
	for _, header := range headers {
		if header.Key == EventTypeIdentifier {
			return Type(header.Value)
		}
	}
	return EventTypeUnknown
}
