package wxsubscript

import "conocourse/config"

// getUsage 返回用法（帮助信息）
func getUsage() string {
	return *config.Usage
}
