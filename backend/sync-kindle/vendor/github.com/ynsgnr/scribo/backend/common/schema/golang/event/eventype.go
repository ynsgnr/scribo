package event

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Type string

const (
	TypeIdentifier = "EventType"

	TypeUnknown            Type = "Unknown"
	TypeAddDevice          Type = "AddDevice"
	TypeAddDeviceSuccess   Type = "AddDeviceSuccess"
	TypeSync2Device        Type = "Sync2Device"
	TypeSync2DeviceSuccess Type = "Sync2DeviceSuccess"
	TypeConvertFile        Type = "ConvertFile"
	TypeConvertFileSuccess Type = "ConvertFileSuccess"
	TypeSend2Device        Type = "Send2Device"
	TypeSend2DeviceSuccess Type = "Send2DeviceSuccess"
	TypeSendMail           Type = "SendMail"
	TypeSendMailSuccess    Type = "SendMailSuccess"
)

// GetEventTypeFromHeaders - returns value for first event type header instance
func GetEventTypeFromHeaders(headers []kafka.Header) Type {
	for _, header := range headers {
		if header.Key == TypeIdentifier {
			return Type(header.Value)
		}
	}
	return TypeUnknown
}
