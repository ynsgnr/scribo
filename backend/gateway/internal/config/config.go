package config

import (
	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	CommandTopic string `env:"COMMAND_TOPIC" default:"command"`

	KafkaEndpoint string `env:"KAFKA" default:"kafka:9092"`
	KafkaGroupID  string `env:"KAFKA_GROUP" default:"myGroup"`
	KafkaOffset   string `env:"KAFKA_OFFSET" default:"earliest"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
