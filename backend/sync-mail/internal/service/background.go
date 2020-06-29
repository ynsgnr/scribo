package service

import (
	"time"

	"github.com/ynsgnr/scribo/backend/common/logger"
)

func (s *service) RunBackground(period time.Duration) {
	defer func() {
		if r := recover(); r != nil {
			logger.Printf(logger.Fatal, "RunBackground: %+v", r)
		}
	}()
	ticker := time.NewTicker(period)
	for {
		select {
		case <-s.done:
			return
		case <-ticker.C:
			err := s.controller.ProcessMails()
			if err != nil {
				logger.Printf(logger.Error, "controller.ApproveKindleMail: %s", err.Error())
			}
		}
	}
}
