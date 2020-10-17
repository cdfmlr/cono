package courseelective

import log "github.com/sirupsen/logrus"

func Init() {
	log.Info("init courseelective")
}

// Run 开启定时从强智教务系统刷新所有学生的所有选课关系的服务。
// 即时刷新一次，然后每周一 05:30 刷新一次
//
// 该方法必须在 {config, model, transport, service/discontinueservice}.Init() 调用完成后执行。
func Run() {
	Refresh()
	CronRefresh()
}
