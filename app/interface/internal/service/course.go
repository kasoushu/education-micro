package service

import (
	"context"
	iv1 "education/api/v1/interface"
	"github.com/go-kratos/kratos/v2/log"
)

type CourseCase interface {
	Create(context.Context, *iv1.CreateCourseReq) error
	Update(context.Context, *iv1.UpdateCourseReq) error
	Delete(context.Context, uint64) error
	GetSingle(context.Context, uint64) (*iv1.CourseInfo, error)
	GetListCourseByTeacherID(context.Context, uint64) (*iv1.CourseListReply, error)
}

func (s *InterfaceService) CreateCourse(ctx context.Context, req *iv1.CreateCourseReq) (*iv1.Reply, error) {
	err := s.courseCase.Create(ctx, req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "create successful!"}, nil
}
func (s *InterfaceService) SaveCourse(ctx context.Context, req *iv1.UpdateCourseReq) (*iv1.Reply, error) {
	err := s.courseCase.Update(ctx, req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "update successful!"}, nil
}
func (s *InterfaceService) DeleteCourse(ctx context.Context, req *iv1.DeleteCourseReq) (*iv1.Reply, error) {
	err := s.courseCase.Delete(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &iv1.Reply{Message: "delete successful!"}, nil
}
func (s *InterfaceService) GetCourse(ctx context.Context, req *iv1.CourseReq) (*iv1.CourseInfo, error) {
	info, err := s.courseCase.GetSingle(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return info, nil
}

func (s *InterfaceService) GetCourseListByTeacherId(ctx context.Context, req *iv1.CourseReq) (*iv1.CourseListReply, error) {
	list, err := s.courseCase.GetListCourseByTeacherID(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return list, nil
}
