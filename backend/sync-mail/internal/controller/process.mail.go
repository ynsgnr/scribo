package controller

import (
	"fmt"

	imap "github.com/emersion/go-imap"
	client "github.com/emersion/go-imap/client"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

func (c *controller) ProcessMails() (err error) {
	//Login to mailer
	imapClient, err := client.DialTLS(fmt.Sprintf("%s:%d", c.imapMailServer, c.imapMailPort), nil)
	if err != nil {
		return
	}
	defer func() {
		logoutErr := imapClient.Logout()
		if logoutErr != nil {
			logger.Printf(logger.Error, "imapClient.Logout: %s", logoutErr.Error())
		}
	}()
	if err = imapClient.Login(c.usernameMail, c.passwordMail); err != nil {
		return
	}

	// Select INBOX
	mbox, err := imapClient.Select(c.imapMailBox, false)
	if err != nil {
		return
	}

	if mbox.Messages == 0 {
		logger.Printf(logger.Info, "ApproveKindleMail: No messages in inbox")
		return nil
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(mbox.Unseen)

	var section imap.BodySectionName
	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- imapClient.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem(), imap.FetchFlags}, messages)
	}()
	for m := range messages {
		//Add new processors here
		err = c.approveMail(m, &section)
		switch err {
		case nil:
		case NotProcessed:
			logger.Printf(logger.Warning, "ProcessMails: %s", err.Error())
		default:
			logger.Printf(logger.Error, "ProcessMails: %s", err.Error())
		}
	}
	err = <-done
	return
}
