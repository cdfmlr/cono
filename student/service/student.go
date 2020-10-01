package service

import (
	"conostudent/model"
	log "github.com/sirupsen/logrus"
)

// GetStudentBySid 获取指定 Sid 的 Student
func GetStudentBySid(sid string) (*model.Student, error) {
	s, err := model.GetStudentBySid(sid)

	logger := log.WithFields(log.Fields{
		"sid":     sid,
		"student": s,
		"err":     err,
	})
	if err != nil {
		logger.Error("GetStudentBySid failed")
	} else {
		logger.Info("GetStudentBySid success")
	}

	return s, err
}

// GetStudentByWechatID 获取指定 WechatID 的 Student
func GetStudentByWechatID(wechatID string) (*model.Student, error) {
	s, err := model.GetStudentByWechatID(wechatID)

	logger := log.WithFields(log.Fields{
		"wechatID": wechatID,
		"student":  s,
		"err":      err,
	})
	if err != nil {
		logger.Error("GetStudentByWechatID failed")
	} else {
		logger.Info("GetStudentByWechatID success")
	}

	return s, err
}

// GetAllStudents 获取全部学生
func GetAllStudents() ([]model.Student, error) {
	s, err := model.GetAllStudents()

	logger := log.WithFields(log.Fields{
		"err": err,
	})
	if err != nil {
		logger.Error("GetAllStudents failed")
	} else {
		logger.Info("GetAllStudents success")
	}

	return s, err
}

// SaveStudent 保存一个学生，存在则更新，不存在则新建
func SaveStudent(student *model.Student) error {
	logger := log.WithFields(log.Fields{"student": student})

	s, err := model.GetStudentBySid(student.Sid)
	if err == nil { // exist
		logger.WithField("origin", s).Info("update student")
		err = s.Update(student)
	} else {
		logger.Info("create student")
		err = model.CreateStudent(student)
	}

	if err != nil {
		logger.WithField("err", err).Error("SaveStudent failed")
	} else {
		logger.Info("SaveStudent success")
	}

	return err
}
