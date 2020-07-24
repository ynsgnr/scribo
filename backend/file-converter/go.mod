module github.com/ynsgnr/scribo/backend/file-converter

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.11
	github.com/confluentinc/confluent-kafka-go v1.4.2
	github.com/golang/protobuf v1.4.2
	github.com/ynsgnr/scribo/backend/common v0.0.0-20200724142247-347f94d54883
	github.com/ynsgnr/scribo/backend/common/configencoding v0.0.0-20200724142247-347f94d54883
	github.com/ynsgnr/scribo/backend/common/logger v0.0.0-20200724142247-347f94d54883
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/ynsgnr/scribo => ../../..
