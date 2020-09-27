package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Student 是学生模型
// 这里指的学生是一个拥有 学号、教务密码，有资格登录教务系统的实体
type Student struct {
	gorm.Model
	Sid      string `gorm:"unique_index:idx_only_one"`
	Password string
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

// UpdatePassword 方法更新一个 Student 的密码
func (s *Student) UpdatePassword(newPassword string) error {
	if s.ID == 0 {
		err := fmt.Errorf("UpdatePassword with an emtpy Student: denied")
		log.WithField("Student", s).Error(err)
		return err
	}
	return DB.Model(&s).Update("password", newPassword).Error
}
