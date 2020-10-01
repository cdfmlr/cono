package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Course 表示一条课程。
//
// 数据条目和 "github.com/cdfmlr/qzgo".GetKbcxAzcRespBodyItem，即「强智教务系统 API —— 课程信息」保持一致。
//
// 参考：
//  - https://github.com/cdfmlr/qzgo/blob/master/getkbcxazc.go
//  - https://qzapi.github.tlingc.com/api/getKbcxAzc/
type Course struct {
	gorm.Model
	Name     string // 课程名称
	Teacher  string // 任课老师
	Location string // 上课地点
	Begin    string // 上课时间
	End      string // 下课时间
	Week     string // 开课周次
	When     string // 上课节次
}

// FindCourse 通过给定参数 condCourse 的非空字段在数据库中查找对应课程（复数）。
// condCourse 不应该包含有被赋值的 gorm.Model 字段，否则不会查询到任何结果。
//
// 如果找到了则返回找到的 Course 们，以及一个 nil；
// 找不到，会返回空 []Course（它的 len == 0），以及一个 nil；
// 有其他任何问题，会返回空 []Course，以及一个非 nil 错误。
func FindCourses(condCourse *Course) ([]Course, error) {
	logger := log.WithFields(log.Fields{
		"condCourse": condCourse,
	})

	var coursesFound []Course

	err := DB.Where(condCourse).Find(&coursesFound).Error

	if err == nil {
		logger.WithField("len_courses_found", len(coursesFound)).Info("FindCourse: success")
	} else {
		logger.WithField("err", err).Error("FindCourse: failed")
	}

	return coursesFound, err
}

// Save 在数据库中保存一条 Course 记录（新建）。
func (c *Course) Save() error {
	logger := log.WithField("course", c)

	err := DB.Where(c).FirstOrCreate(c).Error

	if err != nil {
		logger.WithError(err).Error("Course.Save: failed with unexpected error")
	} else {
		logger.Info("Course.Save: success")
	}

	return err
}

// Delete 从数据库删除一条 Course 记录。
// Course 的 ID 必须不为 0，即 Course 对象不为空，并且必须是从数据库中取出的（例如使用 FindCourse 得到）。
func (c *Course) Delete() error {
	logger := log.WithField("course", *c)

	if c.ID == 0 { // 对象为空
		err := fmt.Errorf("cannot delete a BLANK course")
		logger.WithField("err", err).Warn("Course.Delete: denied")
		return err
	}

	err := DB.Delete(c).Error

	// IMO, wrapping the following lines into a function will make my job easier.
	// However logrus not support a custom Caller depth.
	// There is no way to skip the wrapper function to get a correct ReportCaller.
	// So, I have to write these damn if-else again and again.
	// For updates on this issue:
	// - https://github.com/sirupsen/logrus/issues/972
	// - https://github.com/sirupsen/logrus/pull/989
	if err != nil {
		logger.WithError(err).Error("Course.Delete: failed with unexpected error")
	} else {
		logger.Info("Course.Delete: success")
	}

	return err
}

// Deprecated: use FindCourse instead
// ExistInDB 判断课程是否存在于数据库。
// 判断方法是 Name + Teacher。
// TODO: remove this shit.
func (c *Course) ExistInDB() bool {
	logger := log.WithField("course", *c)

	// 课程名和教师都为空则必不存在
	if c.Name == "" && c.Teacher == "" {
		logger.Info("Course ExistInDB: false")
		return false
	}

	var courseInDB Course

	DB.Where(Course{
		Name:    c.Name,
		Teacher: c.Teacher,
	}).First(&courseInDB)

	exist := courseInDB.ID != 0

	log.WithField("courseInDB_ID", courseInDB.ID).Info("Course ExistInDB:", exist)

	return exist
}
