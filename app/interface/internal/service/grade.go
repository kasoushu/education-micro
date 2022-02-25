package service

import (
	"context"
	iv1 "education/api/v1/interface"
)

type GradeCase interface {
	Create(context.Context, *iv1.GradeReq) error
	UpdateGrade(context.Context, *iv1.GradeUpdateReq) error
	GetGradeByCurriculum(context.Context, *iv1.SingleGradeReq) (*iv1.SingleGradeReply, error)
	GetPeriodListGradeByOneTerm(ctx context.Context, req *iv1.ListPeriodGradeReq) (*iv1.ListGradeReply, error)
	GetGroupListGradeByCurriculum(ctx context.Context, req *iv1.ListGroupGradeReq) (*iv1.ListGradeReply, error)
}

func (s *InterfaceService) SetGrade(ctx context.Context, req *iv1.GradeReq) (*iv1.Reply, error) {
	err := s.gradeCase.Create(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "set successful!"}, nil
}
func (s *InterfaceService) UpdateGrade(ctx context.Context, req *iv1.GradeUpdateReq) (*iv1.Reply, error) {
	err := s.gradeCase.UpdateGrade(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "save successful!"}, nil
}
func (s *InterfaceService) GetGradeByCurriculum(ctx context.Context, req *iv1.SingleGradeReq) (*iv1.SingleGradeReply, error) {
	reply, err := s.gradeCase.GetGradeByCurriculum(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return reply, nil
}
func (s *InterfaceService) GetPeriodListGradeByOneTerm(ctx context.Context, req *iv1.ListPeriodGradeReq) (*iv1.ListGradeReply, error) {
	list, err := s.gradeCase.GetPeriodListGradeByOneTerm(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return list, nil
}
func (s *InterfaceService) GetGroupListGradeByCurriculum(ctx context.Context, req *iv1.ListGroupGradeReq) (*iv1.ListGradeReply, error) {
	list, err := s.gradeCase.GetGroupListGradeByCurriculum(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return list, nil
}
