package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// CoursesBeginHolder 保存所有可能的「开始上课」时间。
// 这些时间相对固定，偶尔刷新一遍就行。
// 用这个可避免每次都遍历DB中所有课程，提高效率。
// TODO: 用 Redis 啊
type CoursesBeginHolder struct {
	coursesBeginTimes []string     // 所有可能的开始上课时间
	ticker            *time.Ticker // coursesBeginTimes 自动刷新的 ticker
	done              chan bool    // 停止自动刷新用的

	sync.RWMutex // coursesBeginTimes 的读写可能并发，用读写锁保证正确
}

// NewCoursesBeginHolder 构造 CoursesBeginHolder 实例
// 构造时即刻刷新一次，然后开始自动刷新，每 refreshPeriod 运行一次 Refresh
func NewCoursesBeginHolder(refreshPeriod time.Duration) *CoursesBeginHolder {
	c := &CoursesBeginHolder{
		coursesBeginTimes: []string{},
		ticker:            time.NewTicker(refreshPeriod),
	}
	c.Refresh()
	c.startAutoRefresh()
	return c
}

// Refresh 刷新 CoursesBeginHolder 的 coursesBeginTimes
func (c *CoursesBeginHolder) Refresh() {
	c.Lock()
	defer c.Unlock()

	// SELECT DISTINCT `begin` FROM `courses` WHERE `courses`.`deleted_at` IS NULL
	model.DB.Model(&model.Course{}).Distinct().Pluck("Begin", &c.coursesBeginTimes)

	log.WithField("coursesBeginTimes", c.coursesBeginTimes).Info("CoursesBeginHolder Refresh")
}

// startAutoRefresh 开始自动刷新
func (c *CoursesBeginHolder) startAutoRefresh() {
	go func() {
		// defer c.refreshTicker.Stop() // 再起不能
		for {
			select {
			case <-c.done:
				log.Info("CoursesBeginHolder AutoRefresh Stop!")
				return
			case <-c.ticker.C:
				c.Refresh()
			}
		}
	}()
}

// GetAll 获取 CoursesBeginHolder 中保存的所有「开始上课」的时间
// 返回值为 []string of "HH:MM"
func (c *CoursesBeginHolder) GetAll() []string {
	c.RLock()
	defer c.RUnlock()

	return c.coursesBeginTimes
}

// GetRecent 获取距离当前最近的一个「开始上课」时间。
// 返回值：
//  - recent time.Time: 这个课程开始的真实的时间，包括完整的年月日时分
//  - cbtStr string: coursesBeginTimes 中的时间字符串表示（"HH:MM"）
func (c *CoursesBeginHolder) GetRecent() (recent time.Time, cbtStr string) {
	// 时区: CST
	loc, _ := time.LoadLocation("PRC")

	// 初始最近的上课时间。
	// FIXME: 2200 年后这个方法不再可用！！
	recent, _ = time.ParseInLocation("Jan 2 2006", "Jan 2 2200", loc)
	recentCbtStr := "25:61"

	c.RLock()
	defer c.RUnlock()

	now := time.Now()

	for _, cbtStr := range c.coursesBeginTimes {
		// string -> time.Time
		cbtHM, _ := time.ParseInLocation("15:04", cbtStr, loc)

		// 加上年月日：要考虑今明两天的这个时间
		today := cbtHM.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
		tomorrow := cbtHM.AddDate(now.Year(), int(now.Month())-1, now.Day())

		// 找出最近的
		for _, cbt := range []time.Time{today, tomorrow} {
			s := cbt.Sub(now)
			if s > 0 && s < recent.Sub(now) {
				recent = cbt
				recentCbtStr = cbtStr
			}
		}
	}

	return recent, recentCbtStr
}

// DefaultCoursesBeginHolder 默认的 CoursesBeginHolder 单例。
// 每 config.CourseNotice.BeginRefreshPeriodSec 秒自动刷新一次。
var DefaultCoursesBeginHolder *CoursesBeginHolder

func initCoursesBeginHolder() {
	log.Info("init service/coursenotice/CoursesBeginHolder: construct DefaultCoursesBeginHolder")

	DefaultCoursesBeginHolder = NewCoursesBeginHolder(time.Duration(config.CourseNotice.BeginRefreshPeriodSec) * time.Second)
}
