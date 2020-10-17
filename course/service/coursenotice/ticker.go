package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
)

// CourseTicker 课程时钟。
// 这东西周期性检查有没有快开始上的课，如果有，则发给选课的学生通知。
//
// TODO(MultiCourseTickers): 考虑同时开多个 CourseTicker 的场景：
//  Tickers 的管理（尤其体现在停止操作上）、打印出来的日志（发生时间重叠时）都将极其混乱！
//  解决这个问题需要考虑：
//  1. 给 CourseTicker 一个 ID;
//  2. 给 CourseTicker 以及调用的下层（如 Notifier ）加上 context 传递。
type CourseTicker struct {
	// Period：运行周期，多久检查一次
	Period time.Duration
	// Threshold：上课阈值，当前时间距离上课小于等于 Threshold 就发起提醒
	Threshold time.Duration

	ticker *time.Ticker
	done   chan bool // 用来停止周期运作的 chan

	coursesBeginHolder *CoursesBeginHolder // 可能的上课时间
	currentWeekHolder  *CurrentWeekHolder  // 当前周次
}

// NewCourseTicker 构造一个 CourseTicker。
// 给定检查周期（period）和上课阈值（threshold）。
//
// 调用该函数，CourseTicker 内部的 time.Ticker 会即刻开始空转，
// 但真正开始要等到 Start 被调用才会开始真正的活动。
func NewCourseTicker(period time.Duration, threshold time.Duration) *CourseTicker {
	c := &CourseTicker{
		Period:             period,
		Threshold:          threshold,
		ticker:             time.NewTicker(period),
		done:               make(chan bool),
		coursesBeginHolder: DefaultCoursesBeginHolder,
		currentWeekHolder:  DefaultCurrentWeekHolder,
	}
	return c
}

// Start 让 CourseTicker 开始周期性工作
func (c *CourseTicker) Start() {
	go func() {
		// 调用时立即执行一次
		c.checkCourses()
		// defer c.ticker.Stop() // 再起不能
		for {
			select {
			case <-c.done:
				log.Info("CourseTicker Stop!")
				return
			case <-c.ticker.C: // 检查上课提醒
				c.checkCourses()
			}
		}
	}()
	// Tip —— Learning Golang:
	// time.Ticker.C 是一个 chan Time，所以可以用 t := <- ticker.C 来获取 tick 发出的时间。
	// 但我们没有使用这个值，而选择在 checkCourses 中用 time.Now() 重新求时间，似乎有些铺张浪费。。
	// 但实际上，这里我是考虑 ticker.C 拿到的时间并不一定是 checkCourses 运行时的时间，
	// 直接把 now := ticker.C 传到 checkCourses 中，可能造成错误的通知！
	// 关于这一点，请看这个 demo：
	//  https://gist.github.com/cdfmlr/a48bfad83cdfbaabc57ceb8d7b60ea81
}

// Stop 停止 CourseTicker
func (c *CourseTicker) Stop() {
	c.done <- true
}

// checkCourses 遍历课程，如果快上课了，就调起通知
func (c *CourseTicker) checkCourses() {
	log.Info("checkCourses...")

	// 最近一个可能上课的时间
	recent, cbtStr := c.coursesBeginHolder.GetRecent()

	// 离上课还早就要不通知
	if recent.Sub(time.Now()) > c.Threshold {
		return
	}

	// 查出星期几、几点开始上的所有课
	courses, err := model.FindCoursesAt(recent.Weekday(), cbtStr)
	if err != nil {
		log.WithError(err).Error("checkCourses failed because of a DB error. Skip this time.")
		return
	}

	// 筛选出是本周要上的课
	currentWeek := c.currentWeekHolder.CurrentWeek()
	ln := 0
	for _, course := range courses {
		if isCourseInWeek(&course, currentWeek) {
			courses[ln] = course
			ln++
		}
	}
	courses = courses[:ln]

	// 通知
	log.WithField("len_courses_to_notify", len(courses)).Info("checkCourses success")

	//go NotifyRecentCourses(courses) // XXX: TODO(MultiCourseTickers) requested
	NotifyRecentCourses(courses)

	log.Info("checkCourses done.")
}

// isCourseInWeek 判断一个 models.Course 是否在指定周次(week) 有课
func isCourseInWeek(course *model.Course, week int) bool {
	parts := strings.Split(course.Week, ",")

	var err error
	var match bool
	for _, w := range parts {
		if match, err = regexp.MatchString(`^(\d*?)-(\d*)$`, w); match {
			// a-b 型
			begin, end := 0, 0
			_, err = fmt.Sscanf(w, "%d-%d", &begin, &end)
			if week >= begin && week <= end {
				return true
			}
		} else if match, err = regexp.MatchString(`^(\d*?)$`, w); match {
			// a 型
			begin := 0
			_, err = fmt.Sscanf(w, "%d", &begin)
			if week == begin {
				return true
			}
		}
	}

	if err != nil {
		log.WithError(err).Error("isCourseInWeek failed")
	}
	return false
}

var DefaultCourseTicker *CourseTicker

func initCourseTicker() {
	log.Info("init CourseTicker: construct DefaultCourseTicker")

	DefaultCourseTicker = NewCourseTicker(
		time.Duration(config.CourseNotice.CoursesCheckPeriodSec)*time.Second,
		time.Duration(config.CourseNotice.RecentCourseThresholdSec)*time.Second)
}
