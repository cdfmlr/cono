package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Student 是学生模型
// 这里指的学生是一个拥有 学号、教务密码，有资格登录教务系统的实体。
// Student 还有一个 WechatID 字段，表示用户在微信公众号中的用户 id
//
// ⚠️注意：数据库不保证 WechatID 的唯一性（也不保证 not null），这些保证由上层调用者来做。
type Student struct {
	gorm.Model
	Sid      string `gorm:"unique;uniqueIndex;not null"`
	Password string `gorm:"not null"`
	WechatID string
}

// GetStudentBySid 获取指定 Sid 的 Student
func GetStudentBySid(sid string) (*Student, error) {
	var student Student
	result := DB.Where("sid = ?", sid).Find(&student)
	if student.ID == 0 {
		log.WithFields(log.Fields{
			"sid":    sid,
			"result": result,
		}).Warn("GetStudentBySid: not exist")
		return &student, fmt.Errorf("student (sid=%s) not exist", sid)
	}
	log.WithFields(log.Fields{
		"student": student,
	}).Info("GetStudentBySid: success")
	return &student, result.Error
}

// GetStudentByWechatID 获取指定 WechatID 的 Student
func GetStudentByWechatID(wechatID string) (*Student, error) {
	var student Student
	result := DB.Where("wechat_id = ?", wechatID).Find(&student)
	if student.ID == 0 {
		log.WithFields(log.Fields{
			"WechatID": wechatID,
			"result":   result,
		}).Warn("GetStudentByWechatID: not exist")
		return &student, fmt.Errorf("student (wechatID=%s) not exist", wechatID)
	}
	log.WithFields(log.Fields{
		"student": student,
	}).Info("GetStudentByWechatID: success")
	return &student, result.Error
}

// CreateStudent 创建一条 student 记录
func CreateStudent(student *Student) error {
	if student.Sid == "" {
		err := fmt.Errorf("denied: create student with empty sid")
		log.Warn("CreateStudent ", err)
		return err
	}

	result := DB.Create(student)

	logger := log.WithFields(log.Fields{
		"student": student,
		"result":  result,
	})
	if result.Error != nil {
		logger.Error("CreateStudent failed")
	} else {
		logger.Info("CreateStudent success")
	}

	return result.Error
}

// GetAllStudents 获取全部学生
func GetAllStudents() ([]Student, error) {
	var students []Student
	result := DB.Find(&students)

	logger := log.WithFields(log.Fields{
		"result": result,
	})
	if result.Error != nil {
		logger.Error("GetAllStudents failed")
	} else {
		logger.Info("GetAllStudents success")
	}

	return students, result.Error
}

// Update 方法更新一个 Student
//    s.Update(newStudent)
//  - s 为原 student，用来在数据库中确定条目（通过其 gorm.Model.ID），xxx指定的非空条目将被更新。
//  - newStudent 是要修改成的信息（不包括 gorm.Model），只有非空字段会被更新。
// GORM 操作参考：https://gorm.io/docs/update.html#Updates-multiple-columns
func (s *Student) Update(newStudent *Student) error {
	if s.ID == 0 {
		err := fmt.Errorf("Update with an emtpy Student: denied")
		log.WithFields(log.Fields{
			"s":   s,
			"new": newStudent,
		}).Error(err)
		return err
	}
	log.WithFields(log.Fields{
		"s":   s,
		"new": newStudent,
	}).Info("Update: success")
	return DB.Model(&s).Updates(newStudent).Error
}
