package service

import (
	"conostudent/config"
	"conostudent/model"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetAllStudents(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	model.Init()
	Init()

	s, err := GetAllStudents()
	if err != nil {
		t.Error("❌ unexpected error:", err)
	}
	t.Logf("len: %v\nitems: %#v", len(s), s)
}

func TestGetStudentBySid(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	model.Init()
	Init()

	// get a exist student
	s, err := GetStudentBySid("000")
	if err != nil {
		t.Error("❌ unexpected", err)
	} else if s.ID == 0 {
		t.Error("❌ unexpected empty result: got s.ID == 0")
	} else {
		t.Log("✅", s)
	}

	// get a NOT EXIST student
	s, err = GetStudentBySid("not_exist")
	if err == nil {
		t.Error("❌ unexpected: get a not_exist, but no err. s =", s)
	} else {
		t.Log("✅", s, err)
	}
}

func TestSave(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	model.Init()
	Init()

	// 已存在的：更新
	err := SaveStudent(&model.Student{
		Sid:      "000",
		Password: "saved_by_TestSave@service/student_test.go",
	})
	if err != nil {
		t.Error("❌ unexpected error:", err)
	} else {
		s, err := GetStudentBySid("000")
		if err != nil {
			t.Error("❌ unexpected error:", err)
		}
		t.Log("👀 s:", s)
	}

	// 不存在的：新建
	randSid := "TestSave" + fmt.Sprint(time.Now().UnixNano()) + fmt.Sprint(rand.Float64())
	err = SaveStudent(&model.Student{
		Sid:      randSid,
		Password: fmt.Sprint(rand.Float64()),
	})
	if err != nil {
		t.Error("❌ unexpected error:", err)
	} else {
		s, err := GetStudentBySid(randSid)
		if err != nil {
			t.Error("❌ unexpected error:", err)
		}
		t.Log("👀 s:", s)
	}
}
