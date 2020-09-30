package model

import (
	"conocourse/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"testing"
	"time"
)

func TestFindCourses(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	// not exist
	cs, err := FindCourses(&Course{
		Name:    "not exist",
		Teacher: "not exist",
	})
	if err != nil {
		t.Error("❌ unexpected err:", err)
	} else if len(cs) != 0 {
		t.Error("❌ unexpected len(cs):", len(cs), "\titems:", cs)
	} else {
		t.Log("✅", len(cs), "\titems:", cs)
	}

	// exist
	cs, err = FindCourses(&Course{
		Name: "TestCourse_Save",
	})
	if err != nil {
		t.Error("❌ unexpected err:", err)
	}
	t.Log("👀", len(cs), cs)
}

func TestCourse_Save(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	c := Course{
		Name:     "TestCourse_Save" + time.Now().Format("15-04-05"),
		Teacher:  "TestCourse",
		Location: "cono/course/model/course_test.go",
		Begin:    time.Now().Format("15:04"),
		End:      "10:00",
		Week:     "1-5",
		When:     "10506",
	}
	err := c.Save()
	if err != nil {
		t.Error("❌ unexpected Save err:", err)
	}
	cc, err := FindCourses(&c)
	if err != nil || len(cc) == 0 {
		t.Log("❌ unexpected err or not found item back:\n--> err:", err, "\n--> cc:", cc)
	}
	t.Log("👀 cc (want only the just saved item): len=", len(cc), "\n--> items:", cc)

}

func TestCourse_Delete(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	c := Course{
		Name:     "TestCourse_Delete" + fmt.Sprint(rand.Int()),
		Teacher:  "TestCourse",
		Location: "cono/course/model/course_test.go",
		Begin:    time.Now().Format("15:04"),
		End:      "10:00",
		Week:     "1-5",
		When:     "10506",
	}

	_ = c.Save()

	// 没主键不能删
	err := c.Delete()
	if err == nil {
		t.Error("❌ unexpected no err")
	} else {
		t.Log("✅", err)
	}
	cc, err := FindCourses(&c)
	if err != nil || len(cc) != 1 {
		t.Error("❌ unexpected : err =", err, "\t len(cc) =", len(cc), "cc =", cc)
	}
	t.Log(cc, err)

	// 下面才应该删了
	_ = cc[0].Delete()
	ccc, err := FindCourses(&c)
	if len(ccc) != 0 || err != nil {
		t.Error("❌ unexpected not deleted: err =", err, "\t ccc =", ccc)
	} else {
		t.Log("✅ blank result and nil error: ", ccc, err)
	}
}
