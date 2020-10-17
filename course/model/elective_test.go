package model

import (
	"conocourse/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestElective_Save(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	var count int64
	DB.Model(&Elective{}).Count(&count)

	type fields struct {
		Model gorm.Model
		Sid   string
		Cid   uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "no exist",
			fields: fields{
				Sid: fmt.Sprint("TestElective_Save", count),
				Cid: uint(count),
			},
			wantErr: false,
		},
		{
			name: "already exist",
			fields: fields{
				Sid: fmt.Sprint("TestElective_Save", count),
				Cid: uint(count),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elective{
				Model: tt.fields.Model,
				Sid:   tt.fields.Sid,
				Cid:   tt.fields.Cid,
			}
			if err := e.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestElective_Delete(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	// Create a record to delete
	var count int64
	DB.Model(&Elective{}).Count(&count)
	toDelete := Elective{
		Sid: fmt.Sprint("TestElective_Delete", count),
		Cid: uint(count),
	}
	if err := toDelete.Save(); err != nil {
		t.Fatal("❌ error to call Elective.Save: err=", err)
	}

	type fields struct {
		Model gorm.Model
		Sid   string
		Cid   uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "delete_not_exist",
			fields: fields{
				Sid: "not_exist",
				Cid: 238985,
			},
			wantErr: true,
		},
		{
			name: "delete_exist",
			fields: fields{
				Model: toDelete.Model,
				Sid:   toDelete.Sid,
				Cid:   toDelete.Cid,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elective{
				Model: tt.fields.Model,
				Sid:   tt.fields.Sid,
				Cid:   tt.fields.Cid,
			}
			if err := e.Delete(); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(e)
		})
	}
}

func TestFindCoursesOfStudent(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	type args struct {
		sid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestFindCoursesOfStudent",
			args:    args{sid: "TestElective_Save"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCourses, err := FindCoursesOfStudent(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCoursesOfStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotCourses) <= 0 {
				t.Errorf("FindCoursesOfStudent() gotCourses (len=%v) = %v", len(gotCourses), gotCourses)
			}
			for i, c := range gotCourses {
				t.Log(i, c)
			}
		})
	}
}

func TestFindStudentsOfCourse(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	type args struct {
		cid uint
	}
	tests := []struct {
		name     string
		args     args
		wantSids []string
		wantErr  bool
	}{
		{
			name:     "TestFindStudentsOfCourse",
			args:     args{cid: 4},
			wantSids: []string{"TestFindStudentsOfCourse1", "TestFindStudentsOfCourse2", "TestFindStudentsOfCourse3"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSids, err := FindStudentsOfCourse(tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindStudentsOfCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSids, tt.wantSids) {
				t.Errorf("FindStudentsOfCourse() gotSids = %v, want %v", gotSids, tt.wantSids)
			}
		})
	}
}

func TestFindElectivesOfCourse(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	var electivesOfCid2 []Elective
	DB.Where("cid = ?", 2).Find(&electivesOfCid2)

	type args struct {
		course Course
	}
	tests := []struct {
		name          string
		args          args
		wantElectives []Elective
		wantErr       bool
	}{
		{
			name:          "empty",
			args:          args{course: Course{}},
			wantElectives: []Elective{},
			wantErr:       false,
		},
		{
			name: "not_exist",
			args: args{course: Course{Model: gorm.Model{
				ID: 99999999999,
			}}},
			wantElectives: []Elective{},
			wantErr:       false,
		},
		{
			name:          "exist",
			args:          args{course: Course{Model: gorm.Model{ID: 2}}},
			wantElectives: electivesOfCid2,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotElectives, err := FindElectivesOfCourse(tt.args.course)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindElectivesOfCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotElectives, tt.wantElectives) {
				t.Errorf("FindElectivesOfCourse() gotElectives = %v, want %v", gotElectives, tt.wantElectives)
			}
			t.Log(gotElectives)
		})
	}
}

func TestFindElectivesOfStudent(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	Init()

	type args struct {
		sid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestFindCoursesOfStudent",
			args:    args{sid: "TestElective_Save"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotElectives, err := FindElectivesOfStudent(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindElectivesOfStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, e := range gotElectives {
				t.Log(i, e.ID, e.Sid, e.Cid, e.Course)
			}
		})
	}
}
