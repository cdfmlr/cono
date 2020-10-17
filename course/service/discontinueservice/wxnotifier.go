package discontinueservice

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/transport"
	log "github.com/sirupsen/logrus"
	"time"
)

// WxDiscontinueServiceNotifier 发送微信公众号模版消息提醒「服务终止」
type WxDiscontinueServiceNotifier struct {
	TemplateMsgID string
	DetailURL     func(student *endpoint.Student, serviceName string, reason string) string
}

// Notify 发微信公众号模版消息通知 student 他的某个服务被停止了。
func (w WxDiscontinueServiceNotifier) Notify(student *endpoint.Student, serviceName string, reason string) {
	// 点击消息打开的详情页面网址
	var detailURL func() string
	if w.DetailURL != nil {
		detailURL = func() string {
			return w.DetailURL(student, serviceName, reason)
		}
	}

	// 构造消息
	msg := endpoint.WxTemplateMessage(
		student.WechatId,
		w.TemplateMsgID,
		detailURL,
		endpoint.WxDiscontinueServiceData(serviceName, time.Now().Format("2006-01-02 15:04:05 MST"), reason))

	log.WithFields(log.Fields{
		"studentSID": student.Sid,
		"msg":        msg,
	}).Info("WxDiscontinueServiceNotifier.Notify")

	// 发送
	transport.WxOfficialAccount.SendTemplateMessage(msg)
}

var DefaultWxDiscontinueServiceNotifier WxDiscontinueServiceNotifier

func initWxDiscontinueServiceNotifier() {
	log.Info("init service/discontinueservice/WxDiscontinueServiceNotifier: construct DefaultWxDiscontinueServiceNotifier")

	DefaultWxDiscontinueServiceNotifier = WxDiscontinueServiceNotifier{
		TemplateMsgID: *config.WxDiscontinueServiceTemplateMsgID,
		DetailURL:     nil,
	}
}
