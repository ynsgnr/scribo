package config

import (
	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	ClientId                string `env:"CLIENT_ID" validate:"required"`
	UserPoolId              string `env:"USER_POOL_ID" validate:"required"`
	InternalGeneratorSecret string `env:"INTERNAL_GENERATOR_SECRET" validate:"required"`
	ExtrenalGeneratorSecret string `env:"EXTERNAL_GENERATOR_SECRET" validate:"required"`

	AddDeviceTopic  string `env:"ADD_DEVICE_TOPIC" default:"device"`
	SyncDeviceTopic string `env:"SYNC_DEVICE_TOPIC" default:"device"`
	EmailTopic      string `env:"EMAIL_TOPIC" default:"email"`

	KafkaEndpoint string `env:"KAFKA" default:"kafka:9092"`
	KafkaGroupID  string `env:"KAFKA_GROUP" default:"myGroup"`
	KafkaOffset   string `env:"KAFKA_OFFSET" default:"earliest"`

	DynamoTableName string `env:"DYNAMO_TABLE_NAME" default:"sync-kindle"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
