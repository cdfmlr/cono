package coursenotice

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/model"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestWxRecentCoursesNotifier_Notify(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	Init()

	courses, _ := model.FindAllCourses()
	testCourse := courses[0]

	testStudent := endpoint.Student{
		Sid:      "not important",
		Password: "useless",
		WechatId: "ow012wl4sXYlxrkCZNQn7_K5iNFQ",
	}

	type fields struct {
		TemplateMsgID string
		DetailURL     func(course *model.Course, student *endpoint.Student) string
	}
	type args struct {
		course  *model.Course
		student *endpoint.Student
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Notify",
			// fields 不定义了，用 DefaultWxRecentCoursesNotifier 来测试
			args: args{
				course:  &testCourse,
				student: &testStudent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := DefaultWxRecentCoursesNotifier
			w.Notify(tt.args.course, tt.args.student)
			<-time.After(time.Second * 5)
		})
	}
}
