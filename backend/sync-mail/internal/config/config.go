package config

import (
	"time"

	"github.com/ynsgnr/scribo/backend/common/configencoding"
)

type Config struct {
	SMTPMailServer string `env:"SMTP_EMAIL" validate:"required"`
	SMTPPort       int    `env:"SMTP_PORT_EMAIL" validate:"required"`
	IMAPMailServer string `env:"IMAP_EMAIL" validate:"required"`
	IMAPPort       int    `env:"IMAP_PORT_EMAIL" validate:"required"`
	IMAPMailBox    string `env:"IMAP_MAIL_BOX_EMAIL" default:"INBOX"`
	UsernameMail   string `env:"USERNAME_EMAIL" validate:"required"`
	PassMail       string `env:"PASS_EMAIL" validate:"required"`
	From           string `env:"FROM_EMAIL" validate:"required"`

	EmailTopic string `env:"EMAIL_TOPIC" default:"email"`

	KafkaEndpoint string `env:"KAFKA" default:"kafka:9092"`
	KafkaGroupID  string `env:"KAFKA_GROUP" default:"sync-mail"`
	KafkaOffset   string `env:"KAFKA_OFFSET" default:"earliest"`

	ApproveKindlePeriod time.Duration `env:"APPROVE_KINDLE_PERIOD" default:"1m"`

	TempFolder string `env:"KAFKA_OFFSET" default:"./tmp"`
	S3Bucket   string `env:"STORAGE_S3_BUCKET" default:"fileconverter"`
}

func InitConfig() (Config, error) {
	cfg := Config{}
	err := configencoding.Set(&cfg)
	return cfg, err
}
