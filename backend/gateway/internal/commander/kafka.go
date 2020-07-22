package commander

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ynsgnr/scribo/backend/common/schema/golang/event"
	"google.golang.org/protobuf/proto"
)

func NewKafkaCommander(producer *kafka.Producer, topic string) Interface {
	return &kafkaCommander{
		producer: producer,
		topic:    topic,
	}
}

type kafkaCommander struct {
	producer *kafka.Producer
	topic    string
}

func (kc kafkaCommander) Send(c Command) error {
	data, ok := c.Data.(proto.Message)
	if !ok {
		return fmt.Errorf("received an interface without proto buf functions")
	}
	msg, err := proto.Marshal(data)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}
	return kc.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kc.topic, Partition: kafka.PartitionAny},
		Key:            []byte(c.Key),
		Value:          msg,
		Headers: []kafka.Header{{
			Key:   string(event.TypeIdentifier),
			Value: []byte(c.EventType),
		}, {
			Key:   string(event.CommandIdentifier),
			Value: []byte(event.GetCommandIDFromHeaders([]kafka.Header{})),
		}},
	}, nil)
}
