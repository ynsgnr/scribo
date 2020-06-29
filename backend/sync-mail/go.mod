module github.com/ynsgnr/scribo/backend/sync-mail

go 1.14

require (
	github.com/aws/aws-sdk-go v1.32.9
	github.com/confluentinc/confluent-kafka-go v1.4.2
	github.com/emersion/go-imap v1.0.5
	github.com/emersion/go-message v0.11.1
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21 // indirect
	github.com/ynsgnr/scribo/backend/common v0.0.0-20200624210845-da940c95b153
	github.com/ynsgnr/scribo/backend/common/configencoding v0.0.0-20200624210845-da940c95b153
	github.com/ynsgnr/scribo/backend/common/logger v0.0.0-20200624210845-da940c95b153
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
