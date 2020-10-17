package coursenotice

import (
	"conocourse/endpoint"
	"conocourse/transport"
	"context"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

// CurrentWeekHolder 保存当前周次
// TODO: 用 Redis 啊
// 调用 NewCurrentWeekHolder 构造时刷新一次，
// 然后每周自动刷新一次。
type CurrentWeekHolder struct {
	currentWeek int // 当前周次

	entryID cron.EntryID // 定时刷新任务的 EntryID
	sync.RWMutex
}

// NewCurrentWeekHolder 构造 CurrentWeekHolder。
// 调用该方法时从教务系统刷新一次，
// 然后每周自动刷新一次。
func NewCurrentWeekHolder() *CurrentWeekHolder {
	c := &CurrentWeekHolder{}
	c.Refresh()
	c.startAutoRefresh()
	return c
}

// CurrentWeek 获取「当前周次」
func (c *CurrentWeekHolder) CurrentWeek() int {
	c.RLock()
	defer c.RUnlock()
	return c.currentWeek
}

// SetCurrentWeek 设置「当前周次」
func (c *CurrentWeekHolder) SetCurrentWeek(currentWeek int) {
	c.Lock()
	defer c.Unlock()
	c.currentWeek = currentWeek
}

// Refresh 通过强智教务系统获取周次
func (c *CurrentWeekHolder) Refresh() {
	// XXX: IMPORTANT 这里的消耗实在太大了，需要在 conostudent 的帮助下重构。

	// 获取所有学生
	stuResp, err := transport.StudentRPCClient.GetAllStudents(context.Background(), &endpoint.Empty{})
	if err != nil {
		log.WithError(err).Error("CurrentWeekHolder Refresh failed: unexpected transport.StudentRPCClient.GetAllStudents error")
	}
	students := stuResp.Students

	// 打乱 students 顺序
	rand.Shuffle(len(students), func(i, j int) {
		students[i], students[j] = students[j], students[i]
	})

	// 逐个尝试登录，获取客户端，成功了就停
	var cli *transport.QzClient
	for _, s := range students {
		cli, err = transport.NewQzClient(s)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
	if err != nil {
		log.WithError(err).Error("CurrentWeekHolder Refresh failed: cannot get a QzClient from any students.")
		return
	}

	log.WithFields(log.Fields{
		"currentWeek_previous":  c.currentWeek,
		"currentWeek_update_to": cli.Current.Zc,
	}).Info("CurrentWeekHolder qzRefresh: refresh success")

	// 设置获取当前周次
	c.Lock()
	defer c.Unlock()
	c.currentWeek = cli.Current.Zc
}

// newRefreshTicker 新建一个自动刷新的 Ticker。
// 调用后会睡眠到下周一，然后每周一 05:25 自动更新。
func (c *CurrentWeekHolder) startAutoRefresh() {
	loc, _ := time.LoadLocation("PRC")

	cc := cron.New(cron.WithLocation(loc))
	entryID, err := cc.AddFunc("25 5 * * MON", func() {
		log.Debug("CurrentWeekHolder refresh cron run")
		c.Refresh()
	})

	if err != nil {
		log.WithError(err).Fatal("CurrentWeekHolder startAutoRefresh failed")
	}

	c.entryID = entryID
}

var DefaultCurrentWeekHolder *CurrentWeekHolder

func initCurrentWeekHolder() {
	log.Info("init service/coursenotice/CurrentWeekHolder: construct DefaultCurrentWeekHolder")

	// Refresh 里用到了随机睡眠
	rand.Seed(time.Now().UnixNano())

	DefaultCurrentWeekHolder = NewCurrentWeekHolder()
}
