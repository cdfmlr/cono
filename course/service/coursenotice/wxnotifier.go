package coursenotice

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/model"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
)

// WxCourseNotifier 发送微信公众号模版消息提醒
type WxRecentCoursesNotifier struct {
	TemplateMsgID string
	DetailURL     func(course *model.Course, student *endpoint.Student) string
}

// Notify 发微信公众号模版消息通知 student 要上 course 课了。
func (w WxRecentCoursesNotifier) Notify(course *model.Course, student *endpoint.Student) {
	// 点击消息打开的详情页面网址
	var detailURL func() string
	if w.DetailURL != nil {
		detailURL = func() string {
			return w.DetailURL(course, student)
		}
	}

	// 构造消息
	msg := endpoint.WxTemplateMessage(
		student.WechatId,
		w.TemplateMsgID,
		detailURL,
		endpoint.WxRecentCoursesNoticeData(course, student))

	log.WithFields(log.Fields{
		"courseID":   course.ID,
		"studentSID": student.Sid,
		"msg":        msg,
	}).Info("WxRecentCoursesNotifier.Notify")

	// 发送
	transport.WxOfficialAccount.SendTemplateMessage(msg)
}

var DefaultWxRecentCoursesNotifier WxRecentCoursesNotifier

func initWxRecentCoursesNotifier() {
	log.Info("init service/coursenotice/WxRecentCoursesNotifier: construct DefaultWxRecentCoursesNotifier")

	DefaultWxRecentCoursesNotifier = WxRecentCoursesNotifier{
		TemplateMsgID: config.CourseNotice.WxRecentCoursesNoticeTemplateMsgID,
		DetailURL:     nil,
	}
}
