package biz

import (
	"context"
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/model"
	"education/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type SelectCase struct {
	selectRepo SelectRepo
	log        *log.Helper
}

type SelectRepo interface {
	Create(context.Context, *model.SelectiveCourse) error
	SetSelective(context.Context, uint64) error
	Delete(context.Context, uint64) error
	GetSingleByStudentId(context.Context, uint64) (*model.SelectInfo, error)
	GetListByCurriculumId(context.Context, uint64) ([]*model.SelectInfo, error)
}

func NewSelectCase(repo SelectRepo, logger log.Logger) service.SelectCase {
	return &SelectCase{
		selectRepo: repo,
		log:        log.NewHelper(log.With(logger, "module", "biz")),
	}
}

func (s *SelectCase) Create(ctx context.Context, req *iv1.CreateSelectReq) error {
	err := s.selectRepo.Create(ctx, &model.SelectiveCourse{
		CurriculumId: req.CurriculumId,
		GroupId:      req.GroupId,
		StudentId:    req.StudentId,
	})
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil

}

func (s *SelectCase) SetSelective(ctx context.Context, id uint64) error {
	err := s.selectRepo.SetSelective(ctx, id)
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *SelectCase) Delete(ctx context.Context, id uint64) error {
	err := s.selectRepo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *SelectCase) GetSingle(ctx context.Context, id uint64) (*iv1.SelectReply, error) {
	info, err := s.selectRepo.GetSingleByStudentId(ctx, id)
	if err != nil {
		s.log.Error(err)
		return nil, nil
	}
	return &iv1.SelectReply{
		Id:             info.Id,
		CurriculumId:   info.CurriculumId,
		GroupId:        info.GroupId,
		StudentId:      info.StudentId,
		CurriculumName: info.CurriculumName,
		GroupName:      info.GroupName,
		StudentName:    info.StudentName,
	}, nil
}

func (s *SelectCase) GetListByCurriculumID(ctx context.Context, id uint64) (*iv1.ListSelectReply, error) {
	list, err := s.selectRepo.GetListByCurriculumId(ctx, id)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	res := make([]*iv1.SelectReply, 0)
	for _, info := range list {
		res = append(res, &iv1.SelectReply{
			Id:             info.Id,
			CurriculumId:   info.CurriculumId,
			GroupId:        info.GroupId,
			StudentId:      info.StudentId,
			CurriculumName: info.CurriculumName,
			GroupName:      info.GroupName,
			StudentName:    info.StudentName,
		})
	}
	return &iv1.ListSelectReply{List: res}, nil
}
