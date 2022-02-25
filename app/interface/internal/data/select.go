package data

import (
	"context"
	"education/app/interface/internal/biz"
	"education/app/interface/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type SelectRepo struct {
	data *Data
	log  *log.Helper
}

func NewSelectRepo(d *Data, logger log.Logger) biz.SelectRepo {
	return &SelectRepo{
		data: d,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (s *SelectRepo) Create(ctx context.Context, sc *model.SelectiveCourse) error {
	cc := model.SelectiveCourse{}

	return nil

}

func (s *SelectRepo) SetSelective(ctx context.Context, id uint64) error {
	cc := model.Curriculum{}

	return nil
}

func (s *SelectRepo) Delete(ctx context.Context, id uint64) error {

	return nil
}

func (s *SelectRepo) GetSingleByStudentId(ctx context.Context, id uint64) (*model.SelectInfo, error) {
	cc := model.SelectInfo{}

	return &cc, nil

}

func (s *SelectRepo) GetListByCurriculumId(ctx context.Context, id uint64) ([]*model.SelectInfo, error) {
	var list []*model.SelectInfo

	return list, nil
}
