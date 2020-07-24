package event

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	uuid "github.com/satori/go.uuid"
)

type CommandID string

const CommandIdentifier = "CommandIdentifier"

// GetCommandIDFromHeaders - returns value for first command identifier header instance
func GetCommandIDFromHeaders(headers []kafka.Header) CommandID {
	for _, header := range headers {
		if header.Key == CommandIdentifier {
			return CommandID(header.Value)
		}
	}
	return CommandID(uuid.NewV4().String())
}
