package data

import "education/app/selectCource/internal/biz"

type CourseRepo struct {
	data *Data
}

func NewCourseRepo(d *Data) biz.CourseRepo {
	return CourseRepo{}
}
