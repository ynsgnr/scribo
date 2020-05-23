package logger

import (
	"bytes"
	"fmt"
	"io"
	"time"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/ynsgnr/scribo/backend/common/logger/vendor/github.com/aws/aws-sdk-go/aws"
)

type writer struct {
	cloudwatch *cloudwatchevents.CloudWatchEvents
	source     string
}

// New - Return a new writer that connects to cloudwatch to write log data async
func New(ses *session.Session, source string) io.Writer {
	svc := cloudwatchevents.New(ses)
	return &logger{
		cloudwatch: svc,
		source:     source,
	}
}

// Write - write function parses and writes given bytes to cloudwatch
func (l *logger) Write(p []byte) (n int, err error) {
	go func() {
		logMessage := bytes.NewBuffer(p).String()
		level, ok := LevelMap[logMessage[0]]
		if !ok{
			level = Log
		}
		_, err := l.cloudwatch.PutEvents(&cloudwatchevents.PutEventsInput{
			Entries: []*cloudwatchevents.PutEventsRequestEntry{
				{
					Detail:     aws.String(fmt.Sprintf("{ \"level\": \"%s\", \"message\": \"%s\" }",level, logMessage[1:])),
					DetailType: aws.String("appLog"),
					Resources: []*string{},
					Source: aws.String(l.source),
					Time: time.Now().UTC(),
				},
			},
		})
		if err != nil {
			fmt.Printf(logMessage)
		}
	}()
	return len(p), nil
}
