package config

import (
	"time"

	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	CommandTopic string `env:"COMMAND_TOPIC" default:"command"`

	KafkaEndpoint string `env:"KAFKA" default:"kafka:9092"`

	CrossOriginAllow            string `env:"CROSS-ORIGIN-ALLOW" default:"http://localhost"`
	CrossOriginAllowCredentials string `env:"CROSS-ORIGIN-ALLOW-CRED" default:"true"`
	CrossOriginAllowMethods     string `env:"CROSS-ORIGIN-ALLOW-METHODS" default:"GET, PUT, POST, DELETE, PATCH, HEAD, OPTIONS"`
	CrossOriginAllowHeaders     string `env:"CROSS-ORIGIN-ALLOW-HEADERS" default:"EventType, Content-Type, Authorization"`
	CrossOriginExposeHeaders    string `env:"CROSS-ORIGIN-EXPOSE-HEADERS" default:"User"`

	BlockerPeriod             time.Duration `env:"EVENT_BLOCKER_PERIOD" default:"100ms"`
	BlockerCleanupPeriod      time.Duration `env:"EVENT_BLOCKER_PERIOD" default:"20s"`
	BlockerThrottleAfterTries int           `env:"EVENT_BLOCKER_THROTTLE_AFTER" default:"7"`
	BlockerMaxWait            time.Duration `env:"EVENT_BLOCKER_MAX_WAIT" default:"2s"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
