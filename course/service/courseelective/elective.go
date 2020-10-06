package courseelective

import (
	"conocourse/model"
	"crypto/md5"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

// Elective 完成学生选课操作。
// courses 中应该是 sid 学生要上的**所有课**。
//
// 该函数会删除 sid 学生之前的全部选课记录。
//
// 如果该函数操作前，数据库中存在一个课程，满足：
//  1. 该课程只与 sid 一个学生有选课关系；
//  2. 该课程不存在于参数 courses 中;
// 则这条课程记录也将惨遭删除。
func Elective(sid string, courses []model.Course) {
	cs := make([]model.Course, len(courses))
	copy(cs, courses)

	clean(sid, cs)
	elective(sid, cs)
}

// clean 删除 sid 学生的全部「过期」选课记录。
//
// 「过期」指的是：一条 Elective of Course 存在于数据库中，但对应的 Course 不在参数 courses 里了。
//
// 如果该函数操作前，数据库中存在一个课程，满足：
//  1. 该课程只与 sid 一个学生有选课关系；
//  2. 该课程不存在于参数 courses 中;
// 则这条课程记录也将惨遭删除。
func clean(sid string, courses []model.Course) {
	logger := log.WithFields(log.Fields{
		"sid":         sid,
		"len_courses": len(courses),
	})

	previousElectives, err := model.FindElectivesOfStudent(sid)
	if err != nil {
		logger.WithError(err).Error("clean failed to find previous Electives of student.")
	}

	// 新课程 courses 中各项的哈希值 map
	newMap := make(map[string]model.Course)
	for _, c := range courses {
		newMap[courseHash(&c)] = c
	}

	for _, e := range previousElectives {
		if _, ok := newMap[courseHash(&e.Course)]; !ok { // 旧的不在新的里了，不需要了
			electivesOfCourse, err := model.FindElectivesOfCourse(e.Course)
			if err != nil {
				log.WithError(err).WithField("course", e.Course).Error("find electives of a previous course failed")
				continue
			}
			// 删除无人问津的课程
			if len(electivesOfCourse) == 1 { // 只有当前这一个学生选了这课，现在也不上这课了，所以把这课删除
				if err := e.Course.Delete(); err != nil {
					log.WithError(err).WithField("course", courses).Error("delete a out-of-date course failed")
				}
			}
			// 删除选课关系
			if err := e.Delete(); err != nil {
				log.WithError(err).WithField("elective", e).Error("delete a out-of-date Elective failed")
			}
		}
	}
}

// courseHash 计算一个课程的哈希
func courseHash(c *model.Course) string {
	// 计算 Name, Teacher, Location, Begin, End, Week, When 的 md5 和
	sl := []string{c.Name, c.Teacher, c.Location, c.Begin, c.End, c.Week, c.When}
	data := []byte(strings.Join(sl, ""))
	return fmt.Sprintf("%x", md5.Sum(data))
}

// elective 保存 courses 中的所有课程，并建立 sid 与这些课程的选课关系
func elective(sid string, courses []model.Course) {
	for _, course := range courses {
		logger := log.WithFields(log.Fields{
			"sid":    sid,
			"course": course,
		})
		// 保存课程
		if err := course.Save(); err != nil {
			logger.WithError(err).Error("elective failed: cannot save course.")
			continue
		}
		// 保存选课关系
		ele := model.Elective{
			Sid: sid,
			Cid: course.ID,
		}
		if err := ele.Save(); err != nil {
			logger.WithError(err).Error("elective failed: cannot save elective (s-c relation).")
			continue
		}
		logger.Info("elective success")
	}
}
