package controller

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
	"github.com/ynsgnr/scribo/backend/sync-mail/internal/storage"
)

type Interface interface {
	SendMail(*mail.SendMail) (*mail.SendMail, error)
	ProcessMails() error
}

func NewController(
	downloadLocation string,
	smtpMailServer string,
	smtpPort int,
	imapMailServer string,
	imapMailPort int,
	imapMailBox string,
	usernameMail string,
	passwordMail string,
	from string,
	storage storage.Interface,
) Interface {
	return &controller{
		downloadLocation: downloadLocation,
		smtpMailServer:   smtpMailServer,
		smtpPort:         smtpPort,
		imapMailServer:   imapMailServer,
		imapMailPort:     imapMailPort,
		imapMailBox:      imapMailBox,
		usernameMail:     usernameMail,
		passwordMail:     passwordMail,
		from:             from,
		storage:          storage,
	}
}

type controller struct {
	downloadLocation string
	smtpMailServer   string
	smtpPort         int
	imapMailServer   string
	imapMailPort     int
	imapMailBox      string
	usernameMail     string
	passwordMail     string
	from             string
	storage          storage.Interface
}
