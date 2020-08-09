package config

import (
	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	CommandTopic string `env:"COMMAND_TOPIC" default:"command"`

	KafkaEndpoint string `env:"KAFKA" default:"kafka:9092"`

	CrossOriginAllow            string `env:"CROSS-ORIGIN-ALLOW" default:"http://localhost"`
	CrossOriginAllowCredentials string `env:"CROSS-ORIGIN-ALLOW-CRED" default:"true"`
	CrossOriginAllowMethods     string `env:"CROSS-ORIGIN-ALLOW-METHODS" default:"GET, PUT, POST, DELETE, PATCH, HEAD"`
	CrossOriginAllowHeaders     string `env:"CROSS-ORIGIN-ALLOW-HEADERS" default:"EventType, Content-Type"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
