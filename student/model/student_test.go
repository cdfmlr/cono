package model

import (
	"conostudent/config"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetStudentBySid(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
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

func TestGetAllStudents(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	Init()

	s, err := GetAllStudents()
	if err != nil {
		t.Error("❌ unexpected error:", err)
	}
	t.Logf("len: %v\nitems: %#v", len(s), s)
}

func TestCreateStudent(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	Init()

	s, _ := GetAllStudents()
	studentNum := len(s)

	// 插入已存在的
	err := CreateStudent(&Student{
		Sid:      "000",
		Password: "bad_one_inserted",
	})
	if err == nil {
		t.Error("❌ create a student with existed sid: no error")
	} else {
		t.Log("✅ expected error:", err)
	}

	// 插入不存在的
	err = CreateStudent(&Student{
		Sid:      "TestCreateStudent" + fmt.Sprint(time.Now().UnixNano()) + fmt.Sprint(rand.Float64()),
		Password: fmt.Sprint(rand.Float64()),
	})
	if err != nil {
		t.Error("❌ create a no-exist student got error:", err)
	}

	s, _ = GetAllStudents()
	if len(s) != studentNum+1 {
		t.Errorf("❌ <len of all students: %v> != <len before oper: %v> + 1", len(s), studentNum)
	}

	t.Logf("✅ len of all students: %v + 1 = %v", studentNum, len(s))
}

func TestStudent_Update(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	Init()

	// update on a exist student
	s, err := GetStudentBySid("000")
	if err != nil {
		t.Error("❌ unexpected:", err)
	}
	err = s.Update(&Student{Password: "new_password_4345"})
	if err != nil {
		t.Error("❌ unexpected:", err)
	}

	// update on a no-exist student
	s = &Student{}
	err = s.Update(&Student{Password: "unexpected_change(TestStudent_UpdatePassword: update on a no-exist student)"})
	if err == nil {
		t.Error("❌ unexpected: no error when updating on a no-exist student")
	} else {
		t.Log("✅ expected error:", err)
	}
}
