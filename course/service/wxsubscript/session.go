package wxsubscript

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Session 表示一条「会话」，即：
// 请求 —> 验证（Verify）-> 具体操作（Continue）
type Session interface {
	generateVerification()
	Verify() string
	Continue(verificationCode string) string
}

// session 作为实现 Session 接口的"抽象类"。
type session struct {
	reqUser      string
	reqContent   string
	verification string
}

func (s *session) generateVerification() {
	randI := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000) // 4位随机数
	randS4 := fmt.Sprintf("%04v", randI)                                   // 4位随机数字字符串
	s.verification = randS4
}

func (s *session) Verify() string {
	panic("implement me")
}

func (s *session) Continue(verificationCode string) string {
	panic("implement me")
}

// isReqVerification 判断请求是否为**验证码**，是则返回 true，否则 false
// 验证码应该是四位随机数字，形如：
//		6982
func isReqVerification(reqContent string) bool {
	matched, _ := regexp.MatchString(`^\d{4}$$`, reqContent)
	return matched
}
