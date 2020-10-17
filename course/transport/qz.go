package transport

import (
	"conocourse/config"
	"conocourse/endpoint"
	"conocourse/model"
	"github.com/cdfmlr/qzgo"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// school 是 jwxt.xxxx.edu.cn 的那个 xxxx。
// 从配置文件中读取
var school string

// QzClient 实现和强智教务系统的通信。
// 提供查课表的 QueryWeekCourses 方法。
//
// 请务必使用 NewQzClient 来构造 QzClient 实例。
//
// 由于底层的 qzgo.Client 登录会不定时过期，
// 所以不宜长时间保持一个 QzClient 对象。
type QzClient struct {
	qzgo.Client

	Auth    *qzgo.AuthUserRespBody
	Current *qzgo.GetCurrentTimeRespBody
}

// NewQzClient 新建一个客户端，
// 完成登录、获取教务时间。
// 参数：
//  - student 是用来登录的学生实例
func NewQzClient(student *endpoint.Student) (*QzClient, error) {
	client := &QzClient{
		Client: qzgo.Client{
			School: school,
			Xh:     student.Sid,
			Pwd:    student.Password,
		},
	}

	// Login
	authResp, err := client.AuthUser()
	client.Auth = authResp
	if err != nil {
		return client, err
	}

	// Query Current jiaowu Time
	currentResp, err := client.GetCurrentTime(time.Now().Format("2006-01-02"))
	client.Current = currentResp

	return client, err
}

// QueryWeekCourses 获取某学生当前周的课程
func (c *QzClient) QueryWeekCourses(student endpoint.Student) ([]model.Course, error) {
	courseResp, err := c.GetKbcxAzc(student.Sid, c.Current.Xnxqh, strconv.Itoa(c.Current.Zc))

	var courses []model.Course
	for _, c := range courseResp {
		courses = append(courses, endpoint.CourseFromQzgo(c))
	}

	return courses, err
}

// initQz 初始化强智教务系统相关的东西
func initQz() {
	log.Info("init Qz")

	school = *config.QzSchool
}
