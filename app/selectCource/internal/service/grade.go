package service

import (
	"context"
	cv1 "education/api/v1/course"
)

type GradeCase interface {
	Create(context.Context, *cv1.GradeReq) error
	UpdateGrade(context.Context, *cv1.GradeUpdateReq) error
	GetGradeByCurriculum(context.Context, *cv1.SingleGradeReq) (*cv1.SingleGradeReply, error)
	GetPeriodListGradeByOneTerm(ctx context.Context, req *cv1.ListPeriodGradeReq) (*cv1.ListGradeReply, error)
	GetGroupListGradeByCurriculum(ctx context.Context, req *cv1.ListGroupGradeReq) (*cv1.ListGradeReply, error)
}

func (s *CourseService) SetGrade(ctx context.Context, req *cv1.GradeReq) (*cv1.Reply, error) {
	err := s.gradeCase.Create(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &cv1.Reply{Message: "set successful!"}, nil
}
func (s *CourseService) UpdateGrade(ctx context.Context, req *cv1.GradeUpdateReq) (*cv1.Reply, error) {
	err := s.gradeCase.UpdateGrade(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &cv1.Reply{Message: "save successful!"}, nil
}
func (s *CourseService) GetGradeByCurriculum(ctx context.Context, req *cv1.SingleGradeReq) (*cv1.SingleGradeReply, error) {
	reply, err := s.gradeCase.GetGradeByCurriculum(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return reply, nil
}
func (s *CourseService) GetPeriodListGradeByOneTerm(ctx context.Context, req *cv1.ListPeriodGradeReq) (*cv1.ListGradeReply, error) {
	list, err := s.gradeCase.GetPeriodListGradeByOneTerm(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return list, nil
}
func (s *CourseService) GetGroupListGradeByCurriculum(ctx context.Context, req *cv1.ListGroupGradeReq) (*cv1.ListGradeReply, error) {
	list, err := s.gradeCase.GetGroupListGradeByCurriculum(ctx, req)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return list, nil
}
