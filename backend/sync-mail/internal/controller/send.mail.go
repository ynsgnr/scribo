package controller

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/ynsgnr/scribo/backend/common/logger"
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/mail"
	"gopkg.in/gomail.v2"
)

func (c *controller) SendMail(sendMail *mail.SendMail) (*mail.SendMail, error) {
	m := gomail.NewMessage()
	from := sendMail.From
	if sendMail.From == "" {
		from = c.from
	}
	m.SetHeader("From", from)
	m.SetHeader("To", sendMail.To)

	err := os.Mkdir(c.downloadLocation, os.ModeDir)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	var filePath string
	if sendMail.AttachmentLocation != "" {
		logger.Printf(logger.Info, "SendMail: to %s : Downloading file: %s", sendMail.To, sendMail.AttachmentLocation)
		filePath, err := c.downloadFile(sendMail.AttachmentLocation, c.downloadLocation)
		if err != nil {
			return nil, err
		}
		m.Attach(filePath)
	} else {
		logger.Printf(logger.Info, "SendMail: to %s : No file to download", sendMail.To)
	}

	err = gomail.NewDialer(c.smtpMailServer, c.smtpPort, c.usernameMail, c.passwordMail).DialAndSend(m)
	if err != nil {
		return nil, err
	}
	if filePath != "" {
		err = os.Remove(filePath)
	}
	return sendMail, err
}

func (c *controller) downloadFile(url string, filepath string) (downloadedFilePath string, err error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	downloadedFilePath = path.Join(filepath, path.Base(r.URL.Path))
	out, err := os.Create(downloadedFilePath)
	if err != nil {
		return
	}
	defer out.Close()
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return
}

func (c *controller) deleteFile(filepath string) error {
	return os.Remove(filepath)
}
