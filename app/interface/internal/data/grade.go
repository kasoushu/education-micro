package data

import (
	"context"
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
	cc := model.Grade{}

	return nil
}

func (g *GradeRepo) Update(ctx context.Context, grade *model.Grade, id uint64) error {
	cc := model.Grade{}

	return nil
}

func (g *GradeRepo) GetGradeByCurriculum(ctx context.Context, GradeQuery *model.GradeQueryByCurriculumOnOneTerm) (*model.SingleGrade, error) {
	cc := model.SingleGrade{}
	return &cc, nil
}

func (g *GradeRepo) GetGradeByStudentId(ctx context.Context, GradeQuery *model.GradeQueryByStudentIdOnOneTerm) ([]*model.SingleGrade, error) {

	var list []*model.SingleGrade

	return list, nil
}

func (g *GradeRepo) GetGradeByGroupId(ctx context.Context, GradeQuery *model.GradeQueryByGroupIdOnOneTerm) ([]*model.SingleGrade, error) {
	var list []*model.SingleGrade

	return list, nil
}
