package controller

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
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
}
