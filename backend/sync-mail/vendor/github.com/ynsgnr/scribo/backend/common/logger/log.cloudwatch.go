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
		now := time.Now().Truncate(time.Microsecond).UTC()
		logMessage := bytes.NewBuffer(p).String()
		msg := logMessage
		level := LevelMap[Default]
		t := now
		var ok bool
		if len(logMessage) >= 21 {
			msg = logMessage[21:]
			msg = strings.ReplaceAll(msg, "\n", " ")
			level, ok = LevelMap[LogLevel(logMessage[20])]
			if !ok {
				level = LevelMap[Default]
				msg = logMessage[20:]
			}
			year, yerr := strconv.Atoi(logMessage[0:4])
			month, merr := strconv.Atoi(logMessage[5:7])
			day, derr := strconv.Atoi(logMessage[8:10])
			hour, herr := strconv.Atoi(logMessage[11:13])
			min, minerr := strconv.Atoi(logMessage[14:16])
			sec, serr := strconv.Atoi(logMessage[17:19])

			if yerr != nil || merr != nil || derr != nil || herr != nil || minerr != nil || serr != nil {
				fmt.Printf("ERROR: parsing time")
			} else {
				t = time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local).UTC()
			}
		}
		timestamp := t.UnixNano() / int64(time.Millisecond)
		message := fmt.Sprintf("(%s) %s: %s", l.source, level, msg)
		//Send the actual request
		l.mutex.Lock()
		logData := &cloudwatchlogs.PutLogEventsInput{
			LogEvents: []*cloudwatchlogs.InputLogEvent{
				{
					Message:   aws.String(message),
					Timestamp: &timestamp,
				},
			},
			LogGroupName:  aws.String(logGroupName),
			LogStreamName: aws.String(logStreamName),
			SequenceToken: seqToken,
		}
		output, err := l.cloudwatch.PutLogEvents(logData)
		if tokenExc, ok := err.(*cloudwatchlogs.InvalidSequenceTokenException); ok {
			if seqToken != nil {
				fmt.Println("WARNING: logging to cloudwatch: ErrCodeInvalidSequenceTokenException")
			}
			logData.SequenceToken = tokenExc.ExpectedSequenceToken
			output, err = l.cloudwatch.PutLogEvents(logData)
		}
		if err != nil {
			fmt.Printf("ERROR: logging to cloudwatch: %s\n", err.Error())
		}
		if output != nil && output.NextSequenceToken != nil {
			seqToken = output.NextSequenceToken
		} else {
			seqToken = nil
		}
		l.mutex.Unlock()
		if output != nil && output.RejectedLogEventsInfo != nil &&
			(output.RejectedLogEventsInfo.ExpiredLogEventEndIndex != nil ||
				output.RejectedLogEventsInfo.TooNewLogEventStartIndex != nil ||
				output.RejectedLogEventsInfo.TooOldLogEventEndIndex != nil) {
			fmt.Printf("ERROR: %s\n", output.GoString())
		}
		fmt.Println(message)
	}()
	return len(p), nil
}
