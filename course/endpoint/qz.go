package endpoint

import (
	"conocourse/model"
	"github.com/cdfmlr/qzgo"
)

// 这个文件里实现和 cdfmlr/qzgo 的接口

// CourseFromQzgo 将一个 qzgo.GetKbcxAzcRespBodyItem 对象转化为 model.Course 对象
func CourseFromQzgo(course qzgo.GetKbcxAzcRespBodyItem) model.Course {
	return model.Course{
		Name:     course.Kcmc,
		Teacher:  course.Jsxm,
		Location: course.Jsmc,
		Begin:    course.Kssj,
		End:      course.Jssj,
		Week:     course.Kkzc,
		When:     course.Kcsj,
	}
}
