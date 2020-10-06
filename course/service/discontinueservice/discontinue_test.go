package discontinueservice

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/model"
	"conocourse/transport"
	"context"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDiscontinueService(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	Init()

	DiscontinueService("201890900999", "测试")
}

func Test_notice(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	transport.Init()
	Init()

	student, err := transport.StudentRPCClient.GetStudentBySid(context.Background(), &endpoint.GetStudentBySidRequest{Sid: "201890900999"})
	if err != nil {
		return
	}
	notice(student, "测试")
}
