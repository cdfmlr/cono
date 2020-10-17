package transport

import (
	"conocourse/config"
	"github.com/cdfmlr/wxofficialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var WxOfficialAccount *wxofficialaccount.WxOfficialAccount

func initWx() {
	log.Info("init WxOfficialAccount")

	WxOfficialAccount = wxofficialaccount.NewWxOfficialAccount(
		config.WxOfficialAccount.AppID,
		config.WxOfficialAccount.AppSecret,
		config.WxOfficialAccount.Token)
}

// WxHttpServe 开启微信公众号的 http 消息接收、回复服务。
func WxHttpServe(handler func(msg message.MixMessage) *message.Reply) {
	WxOfficialAccount.SetMessageHandler(handler)

	log.WithFields(log.Fields{
		"address": config.WxOfficialAccount.Address,
		"handler": handler,
	}).Info("boot WxOfficialAccount Server.")

	err := http.ListenAndServe(config.WxOfficialAccount.Address, WxOfficialAccount)
	if err != nil {
		log.WithField("address", config.WxOfficialAccount.Address).
			WithError(err).
			Fatal("Failed to start WxOfficialAccount Server.")
	}
}
