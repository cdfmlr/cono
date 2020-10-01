package transport

import (
	"conocourse/config"
	"github.com/cdfmlr/wxofficialaccount"
	log "github.com/sirupsen/logrus"
)

var WxOfficialAccount *wxofficialaccount.WxOfficialAccount

func initWx() {
	log.Info("init WxOfficialAccount")

	WxOfficialAccount = wxofficialaccount.NewWxOfficialAccount(
		config.WxOfficialAccount.AppID,
		config.WxOfficialAccount.AppSecret,
		config.WxOfficialAccount.Token)
}
