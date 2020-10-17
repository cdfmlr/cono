package wxsubscript

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"time"
)

const (
	HandleTimeout   = 3 * time.Second
	TimeoutResponse = "处理超时，请稍后重试，若问题仍然存在请联系管理员..."
)

// Handler 处理来自微信公众号 HTTP 服务的请求
func Handler(msg message.MixMessage) *message.Reply {
	var response string

	fromUser := string(msg.FromUserName) // FromUserName 是 CDATA 类型的，这个类型是个 string 的别名

	select {
	case response = <-ResponderInstance().Respond(fromUser, msg.Content):
		break
	case <-time.After(HandleTimeout):
		response = TimeoutResponse
	}

	text := message.NewText(response)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
