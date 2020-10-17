package transport

import (
	"conocourse/config"
	"conocourse/model"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Wx(t *testing.T) {
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	Init()

	// 注意，这里测试的是发客服消息，如果发给不活跃用户可能失败
	// 所以测试前要先让 toUser 随便发条消息给公众号。
	//
	// 客服消息的部分文档：
	//  当用户和公众号产生特定动作的交互时（具体动作列表请见下方说明），微信将会把消息数据推送给开发者，
	//  开发者可以在一段时间内（目前修改为48小时）调用客服接口，通过POST一个JSON数据包来发送消息给普通用户。
	//  此接口主要用于客服等有人工消息处理环节的功能，方便开发者为用户提供更加优质的服务。
	WxOfficialAccount.SendCustomTextMessage("xxx", "你好👋")
}
