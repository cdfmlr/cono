package transport

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/model"
	"context"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Qz(t *testing.T) {
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	Init()

	student, err := StudentRPCClient.GetStudentBySid(
		context.Background(),
		&endpoint.GetStudentBySidRequest{Sid: "201810000999"}) // TODO: fill a real shit.
	if err != nil {
		log.Fatal("StudentRPC Error:", err)
	}

	client, err := NewQzClient(student)
	if err != nil {
		log.Error("❌ unexpected err:", err)
	}

	t.Log("👀 Auth：", client.Auth)
	t.Log("👀 Current：", client.Current)

	cs, err := client.QueryWeekCourses(*student)
	if err != nil {
		log.Error("❌ unexpected err:", err)
	}
	t.Log("👀 Courses：", cs)
}
