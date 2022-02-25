package data

import (
	"context"
	cgv1 "education/api/v1/course"
	"education/app/interface/internal/biz"
	"education/app/interface/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type GradeRepo struct {
	data *Data
	log  *log.Helper
}

func NewGradeRepo(d *Data, logger log.Logger) biz.GradeRepo {
	return &GradeRepo{
		data: d,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (g *GradeRepo) Create(ctx context.Context, grade *model.Grade) error {
	_, err := g.data.courseClient.SetGrade(ctx, &cgv1.GradeReq{
		StudentId:    grade.StudentId,
		CurriculumId: grade.CurriculumId,
		Score:        grade.Score,
		StudyLevel:   grade.StudyLevel,
		GradePoint:   grade.GradePoint,
		Note:         grade.Note,
		GradeYear:    grade.GradeYear,
		Term:         grade.Term,
	})
	if err != nil {
		g.log.Error(err)
		return err
	}
	return nil
}

func (g *GradeRepo) Update(ctx context.Context, grade *model.Grade, id uint64) error {
	_, err := g.data.courseClient.UpdateGrade(ctx, &cgv1.GradeUpdateReq{
		GradeId:      id,
		StudentId:    grade.StudentId,
		CurriculumId: grade.CurriculumId,
		Score:        grade.Score,
		StudyLevel:   grade.StudyLevel,
		GradePoint:   grade.GradePoint,
		Note:         grade.Note,
		GradeYear:    grade.GradeYear,
		Term:         grade.Term,
	})
	if err != nil {
		g.log.Error(err)
		return err
	}
	return nil
}

func (g *GradeRepo) GetGradeByCurriculum(ctx context.Context, GradeQuery *model.GradeQueryByCurriculumOnOneTerm) (*model.SingleGrade, error) {
	v, err := g.data.courseClient.GetGradeByCurriculum(ctx, &cgv1.SingleGradeReq{
		CurriculumId: GradeQuery.CurriculumId,
		StudentId:    GradeQuery.StudentId,
		GradeYear:    GradeQuery.GradeYear,
		Term:         GradeQuery.Term,
	})
	if err != nil {
		g.log.Error(err)
		return nil, err
	}
	return &model.SingleGrade{
		Id:             v.Id,
		CurriculumId:   v.CurriculumId,
		CurriculumName: v.CurriculumName,
		Credit:         v.Credit,
		Category:       v.Category,
		ExamWay:        v.ExamWay,
		StudyLevel:     v.StudyLevel,
		Score:          v.Score,
		GainCredit:     v.GainCredit,
		GradePoint:     v.GradePoint,
		Note:           v.Note,
		StudentName:    v.StudentName,
		StudentId:      v.StudentId,
	}, nil
}

func (g *GradeRepo) GetGradeByStudentId(ctx context.Context, GradeQuery *model.GradeQueryByStudentIdOnOneTerm) ([]*model.SingleGrade, error) {
	var list []*model.SingleGrade
	reply, err := g.data.courseClient.GetPeriodListGradeByOneTerm(ctx, &cgv1.ListPeriodGradeReq{
		StudentId: GradeQuery.StudentId,
		GradeYear: GradeQuery.GradeYear,
		Term:      GradeQuery.Term,
	})
	if err != nil {
		g.log.Error(err)
		return nil, err
	}
	for _, v := range reply.List {
		list = append(list, &model.SingleGrade{
			Id:             v.Id,
			CurriculumId:   v.CurriculumId,
			CurriculumName: v.CurriculumName,
			Credit:         v.Credit,
			Category:       v.Category,
			ExamWay:        v.ExamWay,
			StudyLevel:     v.StudyLevel,
			Score:          v.Score,
			GainCredit:     v.GainCredit,
			GradePoint:     v.GradePoint,
			Note:           v.Note,
			StudentName:    v.StudentName,
			StudentId:      v.StudentId,
		})
	}
	return list, nil
}

func (g *GradeRepo) GetGradeByGroupId(ctx context.Context, GradeQuery *model.GradeQueryByGroupIdOnOneTerm) ([]*model.SingleGrade, error) {
	var list []*model.SingleGrade
	reply, err := g.data.courseClient.GetGroupListGradeByCurriculum(ctx, &cgv1.ListGroupGradeReq{
		GroupId:      GradeQuery.GroupId,
		CurriculumId: GradeQuery.CurriculumId,
		GradeYear:    GradeQuery.GradeYear,
		Term:         GradeQuery.Term,
	})
	if err != nil {
		g.log.Error(err)
		return nil, err
	}
	for _, v := range reply.List {
		list = append(list, &model.SingleGrade{
			Id:             v.Id,
			CurriculumId:   v.CurriculumId,
			CurriculumName: v.CurriculumName,
			Credit:         v.Credit,
			Category:       v.Category,
			ExamWay:        v.ExamWay,
			StudyLevel:     v.StudyLevel,
			Score:          v.Score,
			GainCredit:     v.GainCredit,
			GradePoint:     v.GradePoint,
			Note:           v.Note,
			StudentName:    v.StudentName,
			StudentId:      v.StudentId,
		})
	}
	return list, nil

}
