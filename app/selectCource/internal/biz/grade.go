package biz

import (
	"education/app/selectCource/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type GradeRepo interface {
}

type GradeCase struct {
	gradeRepo GradeRepo
	log       *log.Helper
}

func NewGradeCase(crp GradeRepo, logger log.Logger) service.GradeCase {
	return
}
