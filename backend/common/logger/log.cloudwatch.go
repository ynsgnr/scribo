package logger

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var logGroupName = "beta-scribo-log"
var logStreamName = "backend"
var seqToken *string

type cloudWatchWriter struct {
	cloudwatch *cloudwatchlogs.CloudWatchLogs
	source     string
	mutex      sync.Mutex
}

// New - Return a new writer that connects to cloudwatch to write log data async
func New(ses *session.Session, source string) io.Writer {
	svc := cloudwatchlogs.New(ses)
	return &cloudWatchWriter{
		cloudwatch: svc,
		source:     source,
	}
}

// Write - write function parses and writes given bytes to cloudwatch
func (l *cloudWatchWriter) Write(p []byte) (n int, err error) {
	go func() {
		//Parse log message
		logMessage := bytes.NewBuffer(p).String()
		msg := logMessage
		level := LevelMap[Default]
		t := time.Now().UTC()
		var ok bool
		if len(logMessage) >= 21 {
			msg = logMessage[21:]
			msg = strings.ReplaceAll(msg, "\n", "\\n")
			level, ok = LevelMap[LogLevel(logMessage[20])]
			if !ok {
				level = LevelMap[Default]
			}
			year, yerr := strconv.Atoi(logMessage[0:3])
			month, merr := strconv.Atoi(logMessage[5:6])
			day, derr := strconv.Atoi(logMessage[8:10])
			hour, herr := strconv.Atoi(logMessage[12:13])
			min, minerr := strconv.Atoi(logMessage[15:16])
			sec, serr := strconv.Atoi(logMessage[17:18])

			t = time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)
			if yerr != nil || merr != nil || derr != nil || herr != nil || minerr != nil || serr != nil {
				fmt.Printf("ERROR: parsing time")
				t = time.Now().UTC()
			}
		}
		timestamp := t.Unix()
		//Send the actual request
		l.mutex.Lock()
		output, err := l.cloudwatch.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
			LogEvents: []*cloudwatchlogs.InputLogEvent{
				{
					Message:   aws.String(fmt.Sprintf("%s: %s", level, msg)),
					Timestamp: &timestamp,
				},
			},
			LogGroupName:  aws.String(logGroupName),
			LogStreamName: aws.String(logStreamName),
			SequenceToken: seqToken,
		})
		if err != nil {
			fmt.Printf("ERROR: logging to cloudwatch: %s\n", err.Error())
		}
		if output != nil && output.NextSequenceToken != nil {
			seqToken = output.NextSequenceToken
		} else {
			seqToken = nil
		}
		l.mutex.Unlock()
		fmt.Printf("%s: %s\n", level, msg)
	}()
	return len(p), nil
}
