package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// config is a instance of Conf.
// All configures will be unmarshalled into this var after Init(...) called.
var config Conf

// Exposed configures pointers
var (
	// Database configures
	Database = &config.Database
	// 强智教务系统学校
	QzSchool = &config.QzSchool
	// 微信公众号服务配置
	WxOfficialAccount = &config.WxOfficialAccount
	// StudentRPC 服务地址
	StudentRPCAddress = &config.StudentRPCAddress
	// 课程提醒配置
	CourseNotice = &config.CourseNotice
)

// Init 加载配置文件，写入 Config。
// 加载失败将导致程序 Fatal 退出。
func Init(configFilePath string) {
	logger := log.WithFields(log.Fields{
		"configFilePath": configFilePath,
	})
	logger.Info("init config")

	var err error

	viper.SetConfigFile(configFilePath)

	if err = viper.ReadInConfig(); err != nil {
		logger.WithField("err", err).Fatal("failed to read config file")
	}

	if err = viper.Unmarshal(&config); err != nil {
		logger.WithField("err", err).Fatal("failed to unmarshal config")
	}
}
