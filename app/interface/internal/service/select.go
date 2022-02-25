package service

import (
	"context"
	iv1 "education/api/v1/interface"
)

type SelectCase interface {
	Create(context.Context, *iv1.CreateSelectReq) error
	SetSelective(context.Context, uint64) error
	Delete(context.Context, uint64) error
	GetSingle(context.Context, uint64) (*iv1.SelectReply, error)
	GetListByCurriculumID(context.Context, uint64) (*iv1.ListSelectReply, error)
}

// SetSelective set one curriculum selective
func (s *InterfaceService) SetSelective(ctx context.Context, req *iv1.SetSelectiveReq) (*iv1.Reply, error) {
	err := s.selectCase.SetSelective(ctx, req.CurriculumId)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "set successful!"}, nil
}

// DeleteSelect cancel select
func (s *InterfaceService) DeleteSelect(ctx context.Context, req *iv1.DeleteSelectReq) (*iv1.Reply, error) {
	err := s.selectCase.Delete(ctx, req.Id)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "cancel successful!"}, nil
}

// CreateSelect new select
func (s *InterfaceService) CreateSelect(ctx context.Context, req *iv1.CreateSelectReq) (*iv1.Reply, error) {
	err := s.selectCase.Create(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "select successful!"}, nil
}

// GetSelect Id = studen_id
func (s *InterfaceService) GetSelect(ctx context.Context, req *iv1.GetSelectReq) (*iv1.SelectReply, error) {
	reply, err := s.selectCase.GetSingle(ctx, req.Id)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return reply, nil
}

// GetCourseListSelect get selected list via curriculum_id
func (s *InterfaceService) GetCourseListSelect(ctx context.Context, req *iv1.ListSelectReq) (*iv1.ListSelectReply, error) {
	list, err := s.selectCase.GetListByCurriculumID(ctx, req.CurriculumId)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return list, nil
}
