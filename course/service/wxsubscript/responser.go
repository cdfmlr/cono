package wxsubscript

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

// Responder 对一条消息给出回应
type Responder interface {
	Respond(fromUser string, reqContent string) chan string
}

// responder 实现 Rer。
type responder struct {
	sessions sync.Map // {fromUser: session}
}

func (r *responder) Respond(fromUser string, reqContent string) chan string {
	logger := log.WithFields(log.Fields{
		"fromUser":   fromUser,
		"reqContent": reqContent,
	})

	response := make(chan string)

	// 分发、处理消息，响应放到 response 里
	go func() {
		reqContent = strings.TrimSpace(reqContent) // 去掉首位空白字符

		switch {
		case isReqSubscribe(reqContent):
			logger.Info("respond for Subscribe request")
			s := NewSubscribeSession(fromUser, reqContent)
			r.sessions.Store(fromUser, s)
			response <- s.Verify()
		case isReqUnsubscribe(reqContent):
			logger.Info("respond for UnsubscribeSession request")
			s := NewUnsubscribeSession(fromUser, reqContent)
			r.sessions.Store(fromUser, s)
			response <- s.Verify()
		case isReqVerification(reqContent):
			logger.Info("respond for Verification request")
			if si, loaded := r.sessions.Load(fromUser); loaded {
				s, ok := si.(Session)
				if !ok {
					log.WithFields(log.Fields{
						"si": fmt.Sprintf("%#v", si),
						"s":  s,
						"ok": ok,
					}).Error("Convert interface failed")
					r.sessions.Delete(fromUser)
					response <- "服务器内部错误"
					return
				}
				response <- s.Continue(reqContent)
				r.sessions.Delete(fromUser)
			} else {
				response <- "😯你发这个干嘛？"
			}
		case isReqLicense(reqContent):
			logger.Info("respond for TAC request")
			response <- getLicense()
		default:
			logger.Info("respond usage.")
			response <- getUsage()
		}
	}()

	return response
}

func cleanSessionMap() {
	// TODO: Implement me
	panic("Not yet implemented")
}
