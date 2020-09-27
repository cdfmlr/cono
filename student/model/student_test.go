package model

import (
	"conostudent/config"
	"testing"
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

func TestStudent_UpdatePassword(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	Init()

	// update on a exist student
	s, err := GetStudentBySid("000")
	if err != nil {
		t.Error("❌ unexpected:", err)
	}
	err = s.UpdatePassword("new_password")
	if err != nil {
		t.Error("❌ unexpected:", err)
	}

	// update on a no-exist student
	s = &Student{}
	err = s.UpdatePassword("unexpected_change(TestStudent_UpdatePassword: update on a no-exist student)")
	if err == nil {
		t.Error("❌ unexpected: no error when updating on a no-exist student")
	} else {
		t.Log("✅ expected error:", err)
	}
}
