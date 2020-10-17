package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestCoursenotice_Run(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	Init()

	// Run
	Run()
	<-time.After(time.Second * 10)
}
