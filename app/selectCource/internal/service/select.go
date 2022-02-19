package service

import (
	"context"
	cv1 "education/api/v1/course"
)

type SelectCase interface {
	Create(context.Context, *cv1.CreateSelectReq) error
	SetSelective(context.Context, uint64) error
	Delete(context.Context, uint64) error
	GetSingle(context.Context, uint64) (*cv1.SelectReply, error)
	GetListByCurriculumID(context.Context, uint64) (*cv1.ListSelectReply, error)
}

// SetSelective set one curriculum selective
func (s *CourseService) SetSelective(ctx context.Context, req *cv1.SetSelectiveReq) (*cv1.Reply, error) {
	err := s.selectCase.SetSelective(ctx, req.CurriculumId)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &cv1.Reply{Message: "set successful!"}, nil
}

// DeleteSelect cancel select
func (s *CourseService) DeleteSelect(ctx context.Context, req *cv1.DeleteSelectReq) (*cv1.Reply, error) {
	err := s.selectCase.Delete(ctx, req.Id)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &cv1.Reply{Message: "cancel successful!"}, nil
}

// CreateSelect new select
func (s *CourseService) CreateSelect(ctx context.Context, req *cv1.CreateSelectReq) (*cv1.Reply, error) {
	err := s.selectCase.Create(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &cv1.Reply{Message: "select successful!"}, nil
}

// GetSelect Id = studen_id
func (s *CourseService) GetSelect(ctx context.Context, req *cv1.GetSelectReq) (*cv1.SelectReply, error) {
	reply, err := s.selectCase.GetSingle(ctx, req.Id)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return reply, nil
}

// GetCourseListSelect get selected list via curriculum_id
func (s *CourseService) GetCourseListSelect(ctx context.Context, req *cv1.ListSelectReq) (*cv1.ListSelectReply, error) {
	list, err := s.selectCase.GetListByCurriculumID(ctx, req.CurriculumId)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return list, nil
}
