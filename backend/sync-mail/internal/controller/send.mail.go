package controller

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
)

func (c *controller) SendMail(sendMail *mail.SendMail) (*mail.SendMail, error) {
	from := sendMail.From
	if sendMail.From == "" {
		from = c.from
	}

	e := email.NewEmail()
	e.From = from
	e.To = []string{sendMail.To}

	err := os.Mkdir(c.downloadLocation, os.ModeDir)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	var filePath string
	if sendMail.AttachmentLocation != "" {
		logger.Printf(logger.Info, "SendMail: to %s : Downloading file: %s", sendMail.To, sendMail.AttachmentLocation)
		filePath, err := c.storage.DownloadFile(sendMail.AttachmentLocation)
		if err != nil {
			return nil, err
		}
		e.AttachFile(filePath)
	} else {
		logger.Printf(logger.Info, "SendMail: to %s : No file to download", sendMail.To)
	}

	err = e.Send(fmt.Sprintf("%s:%d", c.smtpMailServer, c.smtpPort), smtp.PlainAuth("", c.usernameMail, c.passwordMail, c.smtpMailServer))
	if err != nil {
		return nil, err
	}
	if filePath != "" {
		err = os.Remove(filePath)
	}
	return sendMail, err
}

func (c *controller) deleteFile(filepath string) error {
	return os.Remove(filepath)
}
