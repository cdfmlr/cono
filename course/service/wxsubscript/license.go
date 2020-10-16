package wxsubscript

import "conocourse/config"

// isReqLicense 判断请求是否为**协议**，即查看《用户协议》和《隐私政策》，是则返回 true，否则 false
// 协议请求应该是：
//		协议
// 这两个字。
func isReqLicense(reqContent string) bool {
	return reqContent == "协议"
}

// getLicense 返回用户协议
func getLicense() string {
	return *config.License
}
