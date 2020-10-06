package courseelective

import (
	"conocourse/endpoint"
	"conocourse/model"
	"conocourse/service/discontinueservice"
	"conocourse/transport"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

const (
	MaxFailedTimes = 3
)

var failedStudents = struct {
	students map[string]uint8 // {sid: failed_times}
	sync.Mutex
}{
	students: map[string]uint8{},
}

// Refresh 从强智教务系统刷新所有学生的所有选课关系。
//
// 注意：密码字段为 DiscontinuedPassword 的学生将不再刷新。
//
// 在每个学生处理完成后，有 0~10 秒的随机睡眠。
// 登录系统失败的学生会被放到 FailedStudents 中待处理。
func Refresh() {
	rand.Seed(time.Now().UnixNano())

	studentsResp, err := transport.StudentRPCClient.GetAllStudents(context.Background(), &endpoint.Empty{})
	if err != nil {
		log.WithError(err).Error("courseelective Refresh failed: transport.StudentRPCClient.GetAllStudents error")
		return
	}
	students := studentsResp.Students
	success := 0

	for _, student := range students {
		if err := refreshStudent(student); err == nil {
			success++
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}

	log.WithFields(log.Fields{
		"success":      success,
		"total":        len(students),
		"success_rate": float32(success) / float32(len(students)),
	}).Info("courseelective Refresh done.")

}

// refreshStudent 从教务系统刷新一个学生的选课
func refreshStudent(student *endpoint.Student) error {
	// 跳过标记为停止服务的学生
	if student.Password == discontinueservice.DiscontinuedPassword {
		log.WithField("student", student).Info("refreshStudent: skip a blank-password student")
		return nil
	}

	// 登录
	qzCli, err := transport.NewQzClient(student)
	if err != nil {
		log.WithError(err).WithField("student", student).Warn("courseelective Refresh: student logging failed. Put student into FailedStudents")

		go handleFailedStudent(student.Sid)
		return err
	}

	// 查课
	kbcxResp, err := qzCli.GetKbcxAzc(student.Sid, qzCli.Current.Xnxqh, fmt.Sprint(qzCli.Current.Zc))

	var courses []model.Course
	for _, k := range kbcxResp {
		courses = append(courses, endpoint.CourseFromQzgo(k))
	}

	// 刷新选课关系
	Elective(student.Sid, courses)

	log.WithFields(log.Fields{
		"student":     student,
		"len_courses": len(courses),
	}).Info("courseelective Refresh student success")

	go handleSuccessStudent(student.Sid)
	return nil
}

// handleFailedStudent 尝试处理刷新失败的学生。
//
// 该函数会在一个小时后重试 refreshStudent。
// 当连续失败次数 > MaxFailedTimes 时，忍痛决定停止对该学生的服务。
func handleFailedStudent(sid string) {
	logger := log.WithField("student", sid)

	failedStudents.Lock()
	defer failedStudents.Unlock()

	failedStudents.students[sid] += 1
	if failedStudents.students[sid] > MaxFailedTimes {
		logger.Warn("handleFailedStudent: failed > MAX_FAILED_TIMES!")
		// 不再服务
		discontinueservice.DiscontinueService(sid, "多次登录教务系统失败，无法获取最新课表。")
		return
	}

	// 一小时后再次尝试
	logger.Info("handleFailedStudent: refresh failed student: Retry in an hour.")
	time.AfterFunc(time.Hour, func() {
		logger.Info("handleFailedStudent: refresh failed student: Retry.")
		student, err := transport.StudentRPCClient.GetStudentBySid(context.Background(), &endpoint.GetStudentBySidRequest{Sid: sid})
		if err != nil {
			logger.WithError(err).Error("handleFailedStudent: get student from conostudent failed!")
			return
		}
		go refreshStudent(student)
	})
}

// handleSuccessStudent 处理刷新成功的学生。
// 清除之前积累的 failed 次数标记。
func handleSuccessStudent(sid string) {
	failedStudents.Lock()
	defer failedStudents.Unlock()

	delete(failedStudents.students, sid)
}

// CronRefresh 创建定时刷新任务。每周一 05:30 调用一次 Refresh。
// 这里考虑用户不会超过 100 人，耗时一般不会超过 1000 秒。所以基本可以在周一第一节课之前更新完成。
func CronRefresh() {
	loc, _ := time.LoadLocation("PRC")

	cc := cron.New(cron.WithLocation(loc))
	_, err := cc.AddFunc("30 5 * * MON", func() {
		log.Debug("courseelective refresh cron run")
		Refresh()
	})

	if err != nil {
		log.WithError(err).Fatal("courseelective add CronRefresh failed")
	}
}
