package biz

import (
	"context"
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/model"
	"education/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type CourseRepo interface {
	Create(context.Context, *model.Curriculum) error
	Update(context.Context, *model.Curriculum, uint64) error
	Delete(context.Context, uint64) error
	GetSingle(context.Context, uint64) (*model.CurriculumItem, error)
	GetListByTeacherId(context.Context, uint64) ([]*model.CurriculumItem, error)
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

func (c *CourseCase) Create(ctx context.Context, req *iv1.CreateCourseReq) error {
	err := c.courseRepo.Create(ctx, &model.Curriculum{
		Name:        req.Name,
		ClassroomId: req.ClassroomId,
		GradeYear:   req.GradeYear,
		Term:        req.Term,
		Credit:      req.Credit,
		SchoolHour:  req.SchoolHour,
		Category:    req.Category,
		Status:      0,
		ExamWay:     req.ExamWay,
		TeacherId:   req.TeacherId,
	})
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *CourseCase) Update(ctx context.Context, req *iv1.UpdateCourseReq) error {

	err := c.courseRepo.Update(ctx, &model.Curriculum{
		Name:        req.Name,
		ClassroomId: req.ClassroomId,
		GradeYear:   req.GradeYear,
		Term:        req.Term,
		Credit:      req.Credit,
		SchoolHour:  req.SchoolHour,
		Category:    req.Category,
		Status:      req.Status,
		ExamWay:     req.ExamWay,
		TeacherId:   req.TeacherId,
	}, req.Id)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil

}

func (c *CourseCase) Delete(ctx context.Context, id uint64) error {
	err := c.courseRepo.Delete(ctx, id)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *CourseCase) GetSingle(ctx context.Context, id uint64) (*iv1.CourseInfo, error) {
	info, err := c.courseRepo.GetSingle(ctx, id)

	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	return &iv1.CourseInfo{
		Id:            info.Id,
		ClassroomId:   info.ClassroomId,
		ClassroomName: info.ClassroomName,
		TeacherId:     info.TeacherId,
		TeacherName:   info.TeacherName,
		Name:          info.Name,
		GradeYear:     info.GradeYear,
		Term:          info.Term,
		Credit:        info.Credit,
		SchoolHour:    info.SchoolHour,
		Category:      info.Category,
		Status:        info.Status,
		ExamWay:       info.ExamWay,
	}, nil

}

func (c *CourseCase) GetListCourseByTeacherID(ctx context.Context, id uint64) (*iv1.CourseListReply, error) {
	list, err := c.courseRepo.GetListByTeacherId(ctx, id)
	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	res := make([]*iv1.CourseInfo, 0)
	for _, info := range list {
		res = append(res, &iv1.CourseInfo{
			Id:            info.Id,
			ClassroomId:   info.ClassroomId,
			ClassroomName: info.ClassroomName,
			TeacherId:     info.TeacherId,
			TeacherName:   info.TeacherName,
			Name:          info.Name,
			GradeYear:     info.GradeYear,
			Term:          info.Term,
			Credit:        info.Credit,
			SchoolHour:    info.SchoolHour,
			Category:      info.Category,
			Status:        info.Status,
			ExamWay:       info.ExamWay,
		})
	}
	return &iv1.CourseListReply{List: res}, nil
}
