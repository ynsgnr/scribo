package controller

import (
	"io"
	"net/http"
	"os"
	"path"

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

	var filePath string
	if sendMail.AttachmentLocation != "" {
		filePath, err := c.downloadFile(sendMail.AttachmentLocation, c.downloadLocation)
		if err != nil {
			return nil, err
		}
		m.Attach(filePath)
	}

	d := gomail.NewDialer(c.smtpMailServer, c.smtpPort, c.usernameMail, c.passwordMail)

	err := d.DialAndSend(m)
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
	out, err := os.Create(filepath)
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
