package wxsubscript

import (
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"sync"
)

/* responder 单例 */
var _responder *responder
var once sync.Once

func ResponderInstance() Responder {
	once.Do(func() {
		_responder = &responder{}
	})
	return _responder
}

func Init() {
	log.Info("init service/wxsubscript...")
	transport.WxOfficialAccount.SetMessageHandler(Handler)
}
