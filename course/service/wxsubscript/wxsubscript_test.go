package wxsubscript

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/service/courseelective"
	"conocourse/service/discontinueservice"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestWxsubscript(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	courseelective.Init()
	discontinueservice.Init()
	Init()

	<-time.After(time.Hour)
}
