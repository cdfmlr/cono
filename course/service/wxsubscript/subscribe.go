package wxsubscript

import (
	"conocourse/endpoint"
	"conocourse/service/courseelective"
	"conocourse/transport"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// isReqSubscribe 判断请求是否为**订阅**操作，是则返回 true，否则 false
// 订阅操作请求内容格式如下：
// 		"订阅课表 201810000999 hd666666"
// 即需符合 "订阅课表" + 空格 + 学号 + 空格 + 教务密码
func isReqSubscribe(reqContent string) bool {
	rs := strings.Split(reqContent, " ")
	if len(rs) == 3 && rs[0] == "订阅课表" { // 符合订阅操作格式
		matched, _ := regexp.MatchString(`^\d{12}$$`, rs[1]) // 学号是数字, 且长度正常
		return matched
	}
	return false
}

// SubscribeSession: 订阅课表的会话
type SubscribeSession struct {
	session

	student endpoint.Student
	//jwClient *transport.QzClient
}

func NewSubscribeSession(reqUser, reqContent string) *SubscribeSession {
	return &SubscribeSession{session: session{
		reqUser:    reqUser,
		reqContent: reqContent,
	}}
}

// Verify 尝试拿用户请求中的信息登录强智系统，检测是否具有办理订阅课表的资格
// 若登录强智系统成功，即用户拥有订阅资格，这是返回强智系统中用户真实姓名、院系、以及一个验证码给用户
//
// 订阅操作请求内容格式如下：
// 		"订阅课表 201810000999 hd666666"
// 即需符合 "订阅课表" + 空格 + 学号 + 空格 + 教务密码
func (s *SubscribeSession) Verify() string {
	rs := strings.Split(s.reqContent, " ")
	sid, pwd := rs[1], rs[2]

	s.student = endpoint.Student{
		Sid:      sid,
		Password: pwd,
		WechatId: s.reqUser,
	}

	// 尝试登录
	qzCli, err := transport.NewQzClient(&s.student)
	//s.jwClient = qzCli

	// 登录失败
	if err != nil {
		log.WithError(err).WithField("student", s.student).Warn("SubscribeSession Verify: student logging failed.")
		return "抱歉，登录教务系统失败，请查正您提供的信息后再试。若问题持续存在，请联系管理员。"
	}

	// 登录成功
	s.generateVerification()
	return fmt.Sprintf(
		"根据您提供的信息，我们查询到您是 %s 的 %s。"+
			"（您的个人信息来自教务系统，仅限验证使用，不会被储存）\n"+
			"如果信息正确无误，且确认订阅课程提醒服务，请回复数字验证码：【%s】(五分钟内有效)",
		qzCli.Auth.UserDwmc,
		qzCli.Auth.UserRealName,
		s.verification,
	)
}

// Continue 为用户办理课程提醒登记，
//  Continue 需要 Verify 提供的验证码
func (s *SubscribeSession) Continue(verificationCode string) string {
	if verificationCode != s.verification { // 验证码错误
		return "验证码错误，以为您取消订阅。"
	}

	log.WithField("student", s.student).Info("SubscribeSession: student subscription added.")

	_ = courseelective.RefreshStudent(&s.student)
	return "订阅成功！\n我们会在每门课上课前通知你哦。🤝"
}
