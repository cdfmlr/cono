package discontinueservice

import log "github.com/sirupsen/logrus"

func Init() {
	log.Info("init service/discontinueservice...")

	initWxDiscontinueServiceNotifier()
}
