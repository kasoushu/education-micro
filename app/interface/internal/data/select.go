package data

import (
	"context"
	cgv1 "education/api/v1/course"
	"education/app/interface/internal/biz"
	"education/app/interface/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type SelectRepo struct {
	data *Data
	log  *log.Helper
}

func NewSelectRepo(d *Data, logger log.Logger) biz.SelectRepo {
	return &SelectRepo{
		data: d,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (s *SelectRepo) Create(ctx context.Context, sc *model.SelectiveCourse) error {
	_, err := s.data.courseClient.CreateSelect(ctx, &cgv1.CreateSelectReq{
		CurriculumId: sc.CurriculumId,
		GroupId:      sc.GroupId,
		StudentId:    sc.StudentId,
	})
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *SelectRepo) SetSelective(ctx context.Context, id uint64) error {
	_, err := s.data.courseClient.SetSelective(ctx, &cgv1.SetSelectiveReq{
		CurriculumId: id,
	})
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *SelectRepo) Delete(ctx context.Context, id uint64) error {
	_, err := s.data.courseClient.DeleteSelect(ctx, &cgv1.DeleteSelectReq{Id: id})
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *SelectRepo) GetSingleByStudentId(ctx context.Context, id uint64) (*model.SelectInfo, error) {
	reply, err := s.data.courseClient.GetSelect(ctx, &cgv1.GetSelectReq{Id: id})
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &model.SelectInfo{
		Id:             reply.Id,
		CurriculumId:   reply.CurriculumId,
		GroupId:        reply.GroupId,
		StudentId:      reply.StudentId,
		StudentName:    reply.StudentName,
		GroupName:      reply.GroupName,
		CurriculumName: reply.CurriculumName,
	}, nil
}

func (s *SelectRepo) GetListByCurriculumId(ctx context.Context, id uint64) ([]*model.SelectInfo, error) {
	var list []*model.SelectInfo
	rep, err := s.data.courseClient.GetCourseListSelect(ctx, &cgv1.ListSelectReq{CurriculumId: id})
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	for _, reply := range rep.List {
		list = append(list, &model.SelectInfo{
			Id:             reply.Id,
			CurriculumId:   reply.CurriculumId,
			GroupId:        reply.GroupId,
			StudentId:      reply.StudentId,
			StudentName:    reply.StudentName,
			GroupName:      reply.GroupName,
			CurriculumName: reply.CurriculumName,
		})
	}
	return list, nil
}
