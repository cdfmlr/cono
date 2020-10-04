package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCurrentWeekHolder(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	Init()

	holder := NewCurrentWeekHolder()
	t.Log(holder.CurrentWeek())

	holder.SetCurrentWeek(666)
	t.Log(holder.CurrentWeek())

	holder.Refresh()
	t.Log(holder.CurrentWeek())
}
