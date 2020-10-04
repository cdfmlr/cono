package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"testing"
)

func TestCourseTicker_refreshCoursesBeginTimes_DryRun(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()

	var times []string
	stmt := model.DB.Session(&gorm.Session{DryRun: true}).
		Model(&model.Course{}).Distinct().Pluck("Begin", &times).Statement

	t.Log(stmt.SQL.String())
}
