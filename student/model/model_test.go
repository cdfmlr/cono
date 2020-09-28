package model

import (
	"conostudent/config"
	"testing"
)

func TestInit(t *testing.T) {
	config.Init("/Users/c/Desktop/StudentConf.yml")
	Init()
	DB.Create(&Student{
		Sid:      "000",
		Password: "test",
	})
	var stus []Student
	DB.Find(&stus)
	t.Log(stus)
}
