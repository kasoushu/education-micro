package data

import (
	"context"
	"education/app/selectCource/internal/biz"
	"education/app/selectCource/internal/model"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	cc := model.Grade{}
	res := g.data.db.WithContext(ctx).Where("student_id=?,curriculum_id=?", grade.StudentId, grade.CurriculumId).First(cc)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.CURRICULUM_HAD_EXISTED
	}
	if res = g.data.db.WithContext(ctx).Create(grade); res.Error != nil {
		g.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (g *GradeRepo) Update(ctx context.Context, grade *model.Grade, id uint64) error {
	cc := model.Grade{}
	res := g.data.db.WithContext(ctx).First(cc, id)
	if res.Error != nil {
		c.log.Error(res.Error)
		return res.Error
	}
	curriculum.Id = id
	res = c.data.db.WithContext(ctx).Model(cc).Updates(curriculum)
	if res.Error != nil {
		c.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (g *GradeRepo) GetGradeByCurriculum(ctx context.Context, GradeQuery *model.GradeQueryByCurriculumOnOneTerm) (*model.SingleGrade, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GradeRepo) GetGradeByStudentId(ctx context.Context, GradeQuery *model.GradeQueryByStudentIdOnOneTerm) ([]*model.SingleGrade, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GradeRepo) GetGradeByGroupId(ctx context.Context, GradeQuery *model.GradeQueryByGroupIdOnOneTerm) ([]*model.SingleGrade, error) {
	//TODO implement me
	panic("implement me")
}
