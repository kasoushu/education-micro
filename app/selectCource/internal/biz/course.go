package biz

import (
	"education/app/selectCource/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type CourseRepo interface {
}

type CourseCase struct {
	courseRepo CourseRepo
	log        *log.Helper
}

func NewCourseCase(crp CourseRepo, logger log.Logger) service.CourseCase {
	return CourseCase{}
}
