package courseelective

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestRefresh(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()

	Refresh()
}
