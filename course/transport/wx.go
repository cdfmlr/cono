package transport

import (
	"conocourse/config"
	"github.com/cdfmlr/wxofficialaccount"
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

	log.WithField("address", config.WxOfficialAccount.Address).Info("boot WxOfficialAccount Server.")
	err := http.ListenAndServe(config.WxOfficialAccount.Address, WxOfficialAccount)
	if err != nil {
		log.WithField("address", config.WxOfficialAccount.Address).
			WithError(err).
			Fatal("Failed to start WxOfficialAccount Server.")
	}
}
