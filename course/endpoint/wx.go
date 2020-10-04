package endpoint

import (
	"conocourse/config"
	"conocourse/model"
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// WxTemplateMessage 帮助构造一条微信公众号消息
func WxTemplateMessage(toUser string, templateID string, detailURL func() string, data map[string]*message.TemplateDataItem) *message.TemplateMessage {
	// 点击消息打开的详情页面网址
	if detailURL == nil {
		// 默认打开教务系统网站
		detailURL = func() string {
			return fmt.Sprintf("http://jwxt.%s.edu.cn", *config.QzSchool)
		}
	}

	// 消息体
	return &message.TemplateMessage{
		ToUser:     toUser,
		TemplateID: templateID,
		URL:        detailURL(),
		Data:       data,
	}
}

// WxRecentCoursesNoticeData 构建快要上课的通知微信消息的 Data (map[string]*message.TemplateDataItem)
//
// emmmm，这个方法随便写的，有点丑。
// 我累了，懒得设计，快点写完了事，问题不大。
func WxRecentCoursesNoticeData(course *model.Course, student *Student) map[string]*message.TemplateDataItem {
	// 猩红字段
	important := func(value string) *message.TemplateDataItem {
		return &message.TemplateDataItem{
			Value: value + "\n",
			Color: "#e51c23",
		}
	}
	// 深蓝字段
	information := func(value string) *message.TemplateDataItem {
		return &message.TemplateDataItem{
			Value: value + "\n",
			Color: "#173177",
		}
	}
	// 漆黑字段
	tip := func(value string) *message.TemplateDataItem {
		return &message.TemplateDataItem{
			Value: value + "\n",
			Color: "#000000",
		}
	}

	return map[string]*message.TemplateDataItem{
		"first":    important("滚去上课!!"),
		"course":   information(course.Name),
		"location": information(course.Location),
		"teacher":  information(course.Teacher),
		"time":     information(course.Begin + "~" + course.End),
		"week":     information(course.Week),
		"bullshit": important("这里留下一句废话的空间。"),
		"remark":   tip("[cono v0.0.0 BETA] Powered By CDFMLR."),
	}
}
