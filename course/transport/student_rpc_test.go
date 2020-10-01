package transport

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/model"
	"context"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_StudentRPC(t *testing.T) {
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	Init()

	resp, err := StudentRPCClient.GetAllStudents(context.Background(), &endpoint.Empty{})
	if err != nil {
		t.Error("❌ unexpected err:", err)
	}
	t.Log(resp)
}
