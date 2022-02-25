package data

import (
	"context"
	cgv1 "education/api/v1/course"
	"education/app/interface/internal/biz"
	"education/app/interface/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type CourseRepo struct {
	data *Data
	log  *log.Helper
}

func NewCourseRepo(d *Data, logger log.Logger) biz.CourseRepo {
	return &CourseRepo{
		data: d,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (c *CourseRepo) Create(ctx context.Context, curriculum *model.Curriculum) error {

	_, err := c.data.courseClient.CreateCourse(ctx, &cgv1.CreateCourseReq{
		ClassroomId: curriculum.ClassroomId,
		TeacherId:   curriculum.TeacherId,
		Name:        curriculum.Name,
		GradeYear:   curriculum.GradeYear,
		Term:        curriculum.Term,
		Credit:      curriculum.Credit,
		SchoolHour:  curriculum.SchoolHour,
		Category:    curriculum.Category,
		ExamWay:     curriculum.ExamWay,
	})
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *CourseRepo) Update(ctx context.Context, curriculum *model.Curriculum, id uint64) error {
	_, err := c.data.courseClient.SaveCourse(ctx, &cgv1.UpdateCourseReq{
		Id:          id,
		ClassroomId: curriculum.ClassroomId,
		TeacherId:   curriculum.TeacherId,
		Name:        curriculum.Name,
		GradeYear:   curriculum.GradeYear,
		Term:        curriculum.Term,
		Credit:      curriculum.Credit,
		SchoolHour:  curriculum.SchoolHour,
		Category:    curriculum.Category,
		Status:      curriculum.Status,
		ExamWay:     curriculum.ExamWay,
	},
	)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *CourseRepo) Delete(ctx context.Context, id uint64) error {
	_, err := c.data.courseClient.DeleteCourse(ctx, &cgv1.DeleteCourseReq{Id: id})
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *CourseRepo) GetSingle(ctx context.Context, id uint64) (*model.CurriculumItem, error) {
	reply, err := c.data.courseClient.GetCourse(ctx, &cgv1.CourseReq{Id: id})

	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	c.log.Debug(reply)
	return &model.CurriculumItem{
		Id:            reply.Id,
		ClassroomId:   reply.ClassroomId,
		ClassroomName: reply.ClassroomName,
		TeacherId:     reply.TeacherId,
		TeacherName:   reply.TeacherName,
		Name:          reply.Name,
		GradeYear:     reply.GradeYear,
		Term:          reply.Term,
		Credit:        reply.Credit,
		SchoolHour:    reply.SchoolHour,
		Category:      reply.Category,
		Status:        reply.Status,
		ExamWay:       reply.ExamWay,
	}, nil
}

func (c *CourseRepo) GetListByTeacherId(ctx context.Context, id uint64) ([]*model.CurriculumItem, error) {
	var list []*model.CurriculumItem
	reply, err := c.data.courseClient.GetCourseListByTeacherId(ctx, &cgv1.CourseReq{Id: id})
	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	for _, v := range reply.List {
		list = append(list, &model.CurriculumItem{
			Id:            v.Id,
			ClassroomId:   v.ClassroomId,
			ClassroomName: v.ClassroomName,
			TeacherId:     v.TeacherId,
			TeacherName:   v.TeacherName,
			Name:          v.Name,
			GradeYear:     v.GradeYear,
			Term:          v.Term,
			Credit:        v.Credit,
			SchoolHour:    v.SchoolHour,
			Category:      v.Category,
			Status:        v.Status,
			ExamWay:       v.ExamWay,
		})
	}

	c.log.Debug(list)
	return list, nil
}
