package courseelective

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/service/discontinueservice"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestRun(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	discontinueservice.Init()
	Init()

	Run()
}
