package endpoint

import (
	"conostudent/model"
)

// StudentFromModel : model.Student -> endpoint.Student
func StudentFromModel(s *model.Student) *Student {
	return &Student{
		Sid:      s.Sid,
		Password: s.Password,
	}
}

// StudentToModel : endpoint.Student -> model.Student
func StudentToModel(s *Student) *model.Student {
	return &model.Student{
		Sid:      s.Sid,
		Password: s.Password,
	}
}
