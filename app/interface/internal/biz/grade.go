package biz

import (
	"context"
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/model"
	"education/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type GradeRepo interface {
	Create(context.Context, *model.Grade) error
	Update(context.Context, *model.Grade, uint64) error
	GetGradeByCurriculum(context.Context, *model.GradeQueryByCurriculumOnOneTerm) (*model.SingleGrade, error)
	GetGradeByStudentId(context.Context, *model.GradeQueryByStudentIdOnOneTerm) ([]*model.SingleGrade, error)
	GetGradeByGroupId(context.Context, *model.GradeQueryByGroupIdOnOneTerm) ([]*model.SingleGrade, error)
}

type GradeCase struct {
	gradeRepo GradeRepo
	log       *log.Helper
}

func (g *GradeCase) SaveGrade(ctx context.Context, req *iv1.GradeUpdateReq) error {
	//TODO implement me
	panic("implement me")
}

func NewGradeCase(crp GradeRepo, logger log.Logger) service.GradeCase {
	return &GradeCase{
		gradeRepo: crp,
		log:       log.NewHelper(log.With(logger, "module", "biz")),
	}
}

func (g *GradeCase) Create(ctx context.Context, req *iv1.GradeReq) error {
	err := g.gradeRepo.Create(ctx, &model.Grade{
		StudentId:    req.StudentId,
		CurriculumId: req.CurriculumId,
		Score:        req.Score,
		StudyLevel:   req.StudyLevel,
		GradePoint:   req.GradePoint,
		Note:         req.Note,
		GradeYear:    req.GradeYear,
		Term:         req.Term,
	})
	if err != nil {
		g.log.Error(err)
		return err
	}
	return nil
}

func (g *GradeCase) UpdateGrade(ctx context.Context, req *iv1.GradeUpdateReq) error {
	err := g.gradeRepo.Update(ctx, &model.Grade{
		StudentId:    req.StudentId,
		CurriculumId: req.CurriculumId,
		Score:        req.Score,
		StudyLevel:   req.StudyLevel,
		GradePoint:   req.GradePoint,
		Note:         req.Note,
		GradeYear:    req.GradeYear,
		Term:         req.Term,
	}, req.GradeId)
	if err != nil {
		g.log.Error(err)
		return err
	}
	return err
}

func (g *GradeCase) GetGradeByCurriculum(ctx context.Context, req *iv1.SingleGradeReq) (*iv1.SingleGradeReply, error) {
	info, err := g.gradeRepo.GetGradeByCurriculum(ctx, &model.GradeQueryByCurriculumOnOneTerm{
		CurriculumId: req.CurriculumId,
		StudentId:    req.StudentId,
		GradeYear:    req.GradeYear,
		Term:         req.Term,
	})
	if err != nil {
		g.log.Error(err)
		return nil, err
	}
	return &iv1.SingleGradeReply{
		Id:             info.Id,
		CurriculumId:   info.CurriculumId,
		CurriculumName: info.CurriculumName,
		Credit:         info.Credit,
		Category:       info.Category,
		ExamWay:        info.ExamWay,
		StudyLevel:     info.StudyLevel,
		Score:          info.Score,
		GainCredit:     info.GainCredit,
		GradePoint:     info.GradePoint,
		Note:           info.Note,
		StudentName:    info.StudentName,
		StudentId:      info.StudentId,
	}, nil
}

func (g *GradeCase) GetPeriodListGradeByOneTerm(ctx context.Context, req *iv1.ListPeriodGradeReq) (*iv1.ListGradeReply, error) {
	list, err := g.gradeRepo.GetGradeByStudentId(ctx, &model.GradeQueryByStudentIdOnOneTerm{
		StudentId: req.StudentId,
		GradeYear: req.GradeYear,
		Term:      req.Term,
	})
	if err != nil {
		g.log.Error(err)
		return nil, err
	}
	res := make([]*iv1.SingleGradeReply, 0)
	for _, info := range list {
		res = append(res, &iv1.SingleGradeReply{
			Id:             info.Id,
			CurriculumId:   info.CurriculumId,
			CurriculumName: info.CurriculumName,
			Credit:         info.Credit,
			Category:       info.Category,
			ExamWay:        info.ExamWay,
			StudyLevel:     info.StudyLevel,
			Score:          info.Score,
			GainCredit:     info.GainCredit,
			GradePoint:     info.GradePoint,
			Note:           info.Note,
			StudentName:    info.StudentName,
			StudentId:      info.StudentId,
		})
	}
	return &iv1.ListGradeReply{List: res}, nil

}

func (g *GradeCase) GetGroupListGradeByCurriculum(ctx context.Context, req *iv1.ListGroupGradeReq) (*iv1.ListGradeReply, error) {
	list, err := g.gradeRepo.GetGradeByGroupId(ctx, &model.GradeQueryByGroupIdOnOneTerm{
		GroupId:      req.GroupId,
		CurriculumId: req.CurriculumId,
		GradeYear:    req.GradeYear,
		Term:         req.Term,
	})
	if err != nil {
		g.log.Error(err)
		return nil, err
	}
	res := make([]*iv1.SingleGradeReply, 0)
	for _, info := range list {
		res = append(res, &iv1.SingleGradeReply{
			Id:             info.Id,
			CurriculumId:   info.CurriculumId,
			CurriculumName: info.CurriculumName,
			Credit:         info.Credit,
			Category:       info.Category,
			ExamWay:        info.ExamWay,
			StudyLevel:     info.StudyLevel,
			Score:          info.Score,
			GainCredit:     info.GainCredit,
			GradePoint:     info.GradePoint,
			Note:           info.Note,
			StudentName:    info.StudentName,
			StudentId:      info.StudentId,
		})
	}
	return &iv1.ListGradeReply{List: res}, nil
}
