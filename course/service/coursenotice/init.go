package coursenotice

import (
	log "github.com/sirupsen/logrus"
)

// Init 初始化 coursenotice 的各个组件。
// 初始化各组件的 Default 实现
func Init() {
	log.Info("init service/coursenotice...")

	initCoursesBeginHolder()
	initCurrentWeekHolder()
	initWxRecentCoursesNotifier()
	initCourseTicker()
}

// Run 启动默认的课程时钟，
// 即调用 DefaultCourseTicker 的 Start 方法，
// 开始课程检查、提醒的定时任务。
//
// 该方法必须在 {config, model, transport, coursenotice}.Init() 调用完成后执行。
func Run() {
	DefaultCourseTicker.Start()
}
