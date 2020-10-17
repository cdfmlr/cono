package wxsubscript

import (
	"conocourse/endpoint"
	"conocourse/service/discontinueservice"
	"conocourse/transport"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// isReqUnsubscribe 判断请求是否为**退订**操作，是则返回 true，否则 false
// 退订操作应该是：
//		退订
// 这两个字。
func isReqUnsubscribe(reqContent string) bool {
	return reqContent == "退订"
}

// Unsubscribe: 退订课表的会话
type UnsubscribeSession struct {
	session
}

func NewUnsubscribeSession(reqUser, reqContent string) *UnsubscribeSession {
	return &UnsubscribeSession{session: session{
		reqUser:    reqUser,
		reqContent: reqContent,
	}}
}

// Verify 尝试退订课程提醒
func (s *UnsubscribeSession) Verify() string {
	s.generateVerification()

	return fmt.Sprintf(
		"您确认要退订课程提醒服务嘛T_T 若您意已决，请回复数字验证码：【%s】(五分钟内有效)",
		s.verification,
	)
}

// Continue 为用户办理课程提醒退订，
//  Continue 需要 Verify 提供的验证码
func (s *UnsubscribeSession) Continue(verificationCode string) string {
	if verificationCode != s.verification { // 验证码错误
		return "验证码错误，以为您取消退订。"
	}

	student, err := transport.StudentRPCClient.GetStudentByWechatID(
		context.Background(),
		&endpoint.GetStudentByWechatIDRequest{WechatId: s.reqUser})

	if err != nil {
		return "退订失败！"
	}

	log.WithField("student", student).Info("UnsubscribeSession: student unsubscribed.")

	discontinueservice.DiscontinueService(student.Sid, "用户退订")
	return "退订成功！"
}
