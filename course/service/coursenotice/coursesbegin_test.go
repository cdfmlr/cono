package coursenotice

import (
	"conocourse/config"
	"conocourse/model"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func TestCoursesBeginHolder_GetAll(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	Init()

	type fields struct {
		coursesBeginTimes []string
		ticker            *time.Ticker
		done              chan bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "GetAll",
			fields: fields{
				coursesBeginTimes: []string{"08:00", "10:30", "12:10"},
				ticker:            nil,
				done:              nil,
			},
			want: []string{"08:00", "10:30", "12:10"},
		},
		{
			name: "empty",
			fields: fields{
				coursesBeginTimes: []string{},
				ticker:            nil,
				done:              nil,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoursesBeginHolder{
				coursesBeginTimes: tt.fields.coursesBeginTimes,
				ticker:            tt.fields.ticker,
				//done:              tt.fields.done,
			}
			if got := c.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestNewCoursesBeginHolder 测试 NewCoursesBeginHolder 的同时，Refresh & startAutoRefresh 也一起测试了
func TestNewCoursesBeginHolder(t *testing.T) {
	// Setup
	log.SetReportCaller(true)
	config.Init("/Users/c/Desktop/CourseConf.yml")
	model.Init()
	Init()

	type args struct {
		refreshPeriod time.Duration
	}
	tests := []struct {
		name string
		args args
		want *CoursesBeginHolder
	}{
		{
			name: "TestNewCoursesBeginHolder",
			args: args{refreshPeriod: time.Second * 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCoursesBeginHolder(tt.args.refreshPeriod)
			t.Log(got)
			<-time.After(time.Second * 10)
		})
	}
}

func TestCoursesBeginHolder_GetRecent(t *testing.T) {
	type fields struct {
		coursesBeginTimes []string
		ticker            *time.Ticker
		done              chan bool
	}
	tests := []struct {
		name       string
		fields     fields
		wantRecent time.Time
		wantCbtStr string
	}{
		{
			name: "GetRecent",
			fields: fields{
				coursesBeginTimes: []string{"08:00", "12:00", "19:00"},
				ticker:            nil,
				done:              nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoursesBeginHolder{
				coursesBeginTimes: tt.fields.coursesBeginTimes,
				//ticker:            tt.fields.ticker,
				//done:              tt.fields.done,
			}
			gotRecent, gotCbtStr := c.GetRecent()
			t.Log("gotRecent:", gotRecent)
			t.Log("gotCbtStr:", gotCbtStr)
		})
	}
}
