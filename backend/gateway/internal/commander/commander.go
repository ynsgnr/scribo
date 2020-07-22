package commander

import "github.com/ynsgnr/scribo/backend/common/schema/golang/event"

type Command struct {
	EventType event.Type
	Key       string
	Data      interface{}
}

type Interface interface {
	Send(Command) error
}
