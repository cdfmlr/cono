package transport

import log "github.com/sirupsen/logrus"

// Init 初始化传输层
func Init() {
	log.Info("init transport...")
	initQz()
	initWx()
	initStudentRPC()
}
