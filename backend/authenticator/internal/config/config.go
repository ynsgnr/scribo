package config

import (
	"time"

	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	ClientId                string `env:"CLIENT_ID" validate:"required"`
	UserPoolId              string `env:"USER_POOL_ID" validate:"required"`
	InternalGeneratorSecret string `env:"INTERNAL_GENERATOR_SECRET" validate:"required"`
	ExtrenalGeneratorSecret string `env:"EXTERNAL_GENERATOR_SECRET" validate:"required"`

	BlockerPeriod             time.Duration `env:"BLOCKER_PERIOD" default:"100ms"`
	BlockerCleanupPeriod      time.Duration `env:"BLOCKER_PERIOD" default:"20s"`
	BlockerThrottleAfterTries int           `env:"BLOCKER_THROTTLE_AFTER" default:"5"`
	BlockerMaxWait            time.Duration `env:"BLOCKER_MAX_WAIT" default:"5s"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
