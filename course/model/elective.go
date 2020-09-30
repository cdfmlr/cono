package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Elective 是选课关系的结构体。
// 用数据库课本上的话来说就是「S-C」表。
type Elective struct {
	gorm.Model

	Sid string // conostudent 服务提供
	Cid uint   // Course.Model.ID

	Course Course `gorm:"foreignKey:Cid"`
}

// Save 保存一条选课关系记录（如果不存在的话）。
//
// 若记录已经存在，则不会做任何操作，也不会返回错误。
// 当且仅当查询、插入数据库时出错才会返回非 nil 错误。
//
// 在操作完成后（如果没有错误发生），调用该方法的对象 e 会更新为数据库中查询到的状态，例如：
//    elective := Elective{Sid: "TestElective_Save2", Cid: 2}
//    // elective is {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}} TestElective_Save2 2}
//    _ = elective.Save()
//    // Now elective is {{4 2020-09-30 15:45:50.075 +0800 CST 2020-09-30 15:45:50.075 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}} TestElective_Save2 2}
func (e *Elective) Save() error {
	logger := log.WithField("elective", e)

	result := DB.Where(e).FirstOrCreate(e)

	if result.Error != nil {
		logger.WithError(result.Error).Error("Elective.Save failed with unexpected error")
	} else {
		logger.Info("Elective.Save success: record saved")
	}

	return result.Error
}

// Delete 删除一条选课关系记录。
// 调用该方法的 Elective 对象 ID 字段必须不为 0，从数据库中直接取出的非空对象可以保证这一点。
func (e *Elective) Delete() error {
	logger := log.WithField("elective", e)

	// XXX: 如果未来需要批量条件删除，(也许)把这个 if 删除就可以。
	//  不过建议保持此方法只允许删除数据库中确实存在的单条记录。
	//  批量删除可以新写一个函数来做，比如 `func DeleteElectives(cond Elective) error`。
	if e.ID == 0 { // 对象为空
		err := fmt.Errorf("cannot delete a BLANK elective")
		logger.WithField("err", err).Warn("Elective.Delete: denied")
		return err
	}

	err := DB.Delete(&e).Error

	if err != nil {
		logger.WithError(err).Error("Elective.Delete: failed with unexpected error")
	} else {
		logger.Info("Elective.Delete: success")
	}

	return err
}

// FindCoursesOfStudent 找出 sid 代表的学生选的所有课。
//
// ⚠️ 注意：如果出错（返回 err != nil 时），返回的 courses 可能为 nil
func FindCoursesOfStudent(sid string) (courses []Course, err error) {
	logger := log.WithField("sid", sid)

	var electives []Elective

	err = DB.Where("sid = ?", sid).Preload("Course").Find(&electives).Error

	if err != nil {
		logger.WithError(err).Error("FindCoursesOfStudent failed: querying electives: got an unexpected error")
		return courses, err
	} else {
		logger.WithField("len_electives_found", len(electives)).Info("FindCoursesOfStudent: querying electives: success")
	}

	courses = []Course{}

	for _, e := range electives {
		courses = append(courses, e.Course)
	}

	return courses, err
}

// FindStudentsOfCourse 找出选了 cid 代表的课程的所有学生的 sid。
//
// ⚠️ 注意：如果出错（返回 err != nil 时），返回的 sids 可能为 nil
func FindStudentsOfCourse(cid uint) (sids []string, err error) {
	logger := log.WithField("cid", cid)

	var electives []Elective

	err = DB.Where("cid = ?", cid).Find(&electives).Error

	if err != nil {
		logger.WithError(err).Error("FindStudentsOfCourse failed: querying electives: got an unexpected error")
		return sids, err
	} else {
		logger.WithField("len_electives_found", len(electives)).Info("FindStudentsOfCourse: querying electives: success")
	}

	sids = []string{}

	for _, e := range electives {
		sids = append(sids, e.Sid)
	}

	return sids, err
}
