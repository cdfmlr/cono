package coursenotice

import (
	"conocourse/endpoint"
	"conocourse/model"
	"conocourse/transport"
	"context"
	log "github.com/sirupsen/logrus"
)

// NotifyRecentCourses 通知上 courses 中每个课的所有学生要上课了
// TODO: 多个课程可以并发处理
func NotifyRecentCourses(courses []model.Course) {
	for _, course := range courses { // 遍历课程
		logger := log.WithField("course", course)
		logger.Info("NotifyRecentCourses...")

		noticed := 0

		// 找出所有选课的学生学号
		sids, err := model.FindStudentsOfCourse(course.ID)
		if err != nil {
			logger.WithError(err).Error("NotifyRecentCourses failed when FindStudentsOfCourse")
			// 如果不为空的话还是通知一下，有一个算一个
			// 是 nil 咱就没法了：
			if sids == nil {
				continue    // nil sids: skip this course. Continue for next course
			}
		}

		for _, sid := range sids { // 遍历选这课的学生
			// 获取学生信息
			student, err := transport.StudentRPCClient.GetStudentBySid(context.Background(),
				&endpoint.GetStudentBySidRequest{Sid: sid})
			if err != nil {
				logger.WithError(err).Error("NotifyRecentCourses failed when GetStudentBySid")
				continue
			}
			// 通知
			DefaultWxRecentCoursesNotifier.Notify(&course, student)
			noticed++
		}

		logger.WithField("noticed_students", noticed).Info("Notify Recent Courses")
	}
}

type RecentCoursesNotifier interface {
	Notify(course *model.Course, student *endpoint.Student)
}
