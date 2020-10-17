package service

import (
	"conocourse/service/courseelective"
	"conocourse/service/coursenotice"
	"conocourse/service/discontinueservice"
	"conocourse/service/wxsubscript"
	log "github.com/sirupsen/logrus"
)

// Init 初始化各种服务
func Init() {
	log.Info("init service")

	discontinueservice.Init()
	courseelective.Init()
	coursenotice.Init()
	wxsubscript.Init()
}

// Run 开启各种服务：
//  - 微信公众号消息服务
//  - 选课自动刷新
//  - 上课时间提醒
func Run() {
	courseelective.Run()
	coursenotice.Run()
	wxsubscript.Run()
}
