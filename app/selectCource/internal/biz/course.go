package biz

import (
	"context"
	cv1 "education/api/v1/course"
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
	return &CourseCase{
		courseRepo: crp,
		log:        log.NewHelper(log.With(logger, "module", "biz")),
	}
}

func (c *CourseCase) Create(ctx context.Context, req *cv1.CreateCourseReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseCase) Update(ctx context.Context, req *cv1.UpdateCourseReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseCase) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseCase) GetSingle(ctx context.Context, id uint64) (*cv1.CourseInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseCase) GetListCourseByTeacherID(ctx context.Context, id uint64) (*cv1.CourseListReply, error) {
	//TODO implement me
	panic("implement me")
}
