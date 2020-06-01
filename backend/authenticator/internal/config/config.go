package config

import (
	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	ClientId   string `env:"CLIENT_ID" validate:"required"`
	UserPoolId string `env:"USER_POOL_ID" validate:"required"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
