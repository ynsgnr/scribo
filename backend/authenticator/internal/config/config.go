package config

import (
	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	ClientId   string `env:"CLIENTID" validate:"required"`
	UserPoolId string `env:"USERPOOLID" validate:"required"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
