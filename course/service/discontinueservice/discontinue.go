package discontinueservice

import (
	"conocourse/endpoint"
	"conocourse/model"
	"conocourse/transport"
	"context"
	log "github.com/sirupsen/logrus"
)

// DiscontinuedPassword 不再服务的学生密码会被写成这个字符串
const DiscontinuedPassword = "<cono> DISCONTINUED </cono>"

// DiscontinueService 停止对 sid 学生的服务
//
// 发微信公众号模版消息提醒、删库。
// 不可撤销！
func DiscontinueService(sid string, reason string) {
	log.WithFields(log.Fields{"sid": sid, "reason": reason}).Debug("DiscontinueService run.")

	student, err := transport.StudentRPCClient.GetStudentBySid(context.Background(), &endpoint.GetStudentBySidRequest{Sid: sid})
	if err != nil {
		return
	}

	notice(student, reason)
	deleteStudent(student)
}

// notice 发微信公众号模版消息告诉 student 不再服务他了。
func notice(student *endpoint.Student, reason string) {
	DefaultWxDiscontinueServiceNotifier.Notify(student, "上课提醒", reason)
}

func deleteStudent(student *endpoint.Student) {
	// 删除选课关系
	electives, err := model.FindElectivesOfStudent(student.Sid)
	if err != nil {
		log.WithError(err).WithField("student", student).Error("find electives of student failed")
		return
	}
	for _, e := range electives {
		if err := e.Delete(); err != nil {
			log.WithError(err).WithField("elective", e).Error("delete Elective failed")
		}
	}
	// "删除"学生
	student.Password = DiscontinuedPassword
	_, err = transport.StudentRPCClient.Save(context.Background(), student)
	if err != nil {
		log.WithError(err).WithField("student", student).Error("save student failed")
	}
}
