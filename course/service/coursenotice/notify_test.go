package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNotifyRecentCourses(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	Init()

	courses, _ := model.FindCoursesOfStudent("201890900999")

	NotifyRecentCourses(courses)
}
