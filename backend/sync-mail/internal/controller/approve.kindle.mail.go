package controller

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-message/mail"
)

var NotProcessed = errors.New("mail is not processed")

const (
	mailBoxName = "do-not-reply"
	hostName    = "amazon.com"
)

func (c *controller) approveMail(message *imap.Message, section *imap.BodySectionName) error {
	isProcessed := false
	for _, from := range message.Envelope.From {
		if from.MailboxName == mailBoxName && from.HostName == hostName {
			r := message.GetBody(section)
			mr, err := mail.CreateReader(r)
			if err != nil {
				return fmt.Errorf("approveMail: mail.CreateReader: %w", err)
			} // Process each message's part
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					return fmt.Errorf("approveMail: mail.Body.NextPart: %w", err)
				}
				switch p.Header.(type) {
				case *mail.InlineHeader:
					// This is the message's text (can be plain-text or HTML)
					b, err := ioutil.ReadAll(p.Body)
					if err != nil {
						return fmt.Errorf("approveMail: mail.Body.NextPart: ReadAll: %w", err)
					}
					body := string(b)
					splitted := strings.Split(body, "<h3 class=\"verification-email-h3\">")
					if len(splitted) < 2 {
						return fmt.Errorf("approveMail: mail.Body.NextPart: Parsing: h3 title parsing doesn't have enough length")
					}
					splitted = strings.Split(splitted[1], "<a href=\"")
					if len(splitted) < 2 {
						return fmt.Errorf("approveMail: mail.Body.NextPart: Parsing: href link parsing doesn't have enough length")
					}
					splitted = strings.Split(splitted[1], "\" class=\"button\"")
					if len(splitted) < 1 {
						return fmt.Errorf("approveMail: mail.Body.NextPart: Parsing: button link parsing doesn't have enough length")
					}
					link := splitted[0]
					_, err = http.Get(link)
					if err != nil {
						return fmt.Errorf("approveMail: mail.Body.NextPart: Get Link: http.Get: %w", err)
					}
					isProcessed = true
				}
			}
		}
	}
	if isProcessed {
		return nil
	} else {
		return NotProcessed
	}
}
