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

// templateDataItemStyle 是公众号模版消息字段风格化函数。
// 具体的实现有下面的 important, information 和 tip。
type templateDataItemStyle func(value string) *message.TemplateDataItem

// important 公众号模版消息——猩红字段
func important(value string) *message.TemplateDataItem {
	return &message.TemplateDataItem{
		Value: value + "\n",
		Color: "#e51c23",
	}
}

// information 公众号模版消息——深蓝字段
func information(value string) *message.TemplateDataItem {
	return &message.TemplateDataItem{
		Value: value + "\n",
		Color: "#173177",
	}
}

// tip 公众号模版消息——漆黑字段
func tip(value string) *message.TemplateDataItem {
	return &message.TemplateDataItem{
		Value: value + "\n",
		Color: "#000000",
	}
}

// WxRecentCoursesNoticeData 构建快要上课的通知微信消息的 Data (map[string]*message.TemplateDataItem)
//
// 对应的模版如下：
//    {{first.DATA}}
//    课程：{{course.DATA}}
//    地点：{{location.DATA}}
//    老师：{{teacher.DATA}}
//    时间：{{time.DATA}}
//    教学周：{{week.DATA}}
//    ---
//    {{bullshit.DATA}}
//    {{remark.DATA}}
func WxRecentCoursesNoticeData(course *model.Course, student *Student) map[string]*message.TemplateDataItem {
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

// WxDiscontinueServiceData 构建「终止服务提醒」的模版消息的 Data (map[string]*message.TemplateDataItem)
//
// 对应的模版如下：
//    {{first.DATA}}
//    终止服务：{{service.DATA}}
//    生效时间：{{time.DATA}}
//    原因：{{reason.DATA}}
//    {{remark.DATA}}
func WxDiscontinueServiceData(serviceName string, vtime string, reason string) map[string]*message.TemplateDataItem {
	return map[string]*message.TemplateDataItem{
		"first":   important("抱歉，cono 即将终止您订阅的某些服务。\n对此，如有任何建议和意见，敬请反馈。期待您的再次使用。"),
		"service": information(serviceName),
		"time":    information(vtime),
		"reason":  information(reason),
		"remark":  tip("解释权归 cono 所有。如有疑问，敬请致电垂询：110。"),
	}
}
