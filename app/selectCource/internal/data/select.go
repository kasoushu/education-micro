package data

import (
	"context"
	"education/app/selectCource/internal/biz"
	"education/app/selectCource/internal/model"
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
	//TODO implement me
	panic("implement me")
}

func (s *SelectRepo) SetSelective(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (s *SelectRepo) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (s *SelectRepo) GetSingleByStudentId(ctx context.Context, id uint64) (*model.SelectInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SelectRepo) GetListByCurriculumId(ctx context.Context, id uint64) ([]*model.SelectInfo, error) {
	//TODO implement me
	panic("implement me")
}
