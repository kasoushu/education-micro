package data

import (
	"context"
	"education/app/selectCource/internal/biz"
	"education/app/selectCource/internal/model"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	res := s.data.db.WithContext(ctx).First(cc, "student_id=?", sc.StudentId)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.CURRICULUM_HAD_EXISTED
	}
	if res = s.data.db.WithContext(ctx).Create(sc); res.Error != nil {
		s.log.Error(res.Error)
		return res.Error
	}
	return nil

}

func (s *SelectRepo) SetSelective(ctx context.Context, id uint64) error {
	cc := model.Curriculum{}
	res := s.data.db.WithContext(ctx).First(&cc, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.CURRICULUM_NOT_FOUND
	}
	res = s.data.db.WithContext(ctx).Model(cc).Update("status", 1)
	//s.data.db.WithContext(ctx).Table("curriculums").Update("status", 1)
	if res.Error != nil {
		s.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (s *SelectRepo) Delete(ctx context.Context, id uint64) error {
	res := s.data.db.WithContext(ctx).Delete(&model.Curriculum{}, id)
	if res.Error != nil {
		s.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (s *SelectRepo) GetSingleByStudentId(ctx context.Context, id uint64) (*model.SelectInfo, error) {
	cc := model.SelectInfo{}
	res := s.data.db.Table("selective_courses").
		Select(`selective_courses.id as id,
						users.id as student_id,
						curriculums.name as name,
						curriculums.id as curriculum_id,
						users.name as student_name,
						groups.id as group_id,
						groups.name as group_name
							`).
		Joins("inner join groups on groups.id=selective_courses.group_id ").
		Joins("inner join users on selective_courses.student_id=users.id and users.is_student=?", true).
		Where("users.id=?", id).First(&cc)
	s.log.Debug(cc)
	if res.Error != nil {
		s.log.Error(res.Error)
		return nil, res.Error
	}
	return &cc, nil

}

func (s *SelectRepo) GetListByCurriculumId(ctx context.Context, id uint64) ([]*model.SelectInfo, error) {
	var list []*model.SelectInfo
	res := s.data.db.Table("selective_courses").
		Select(`selective_courses.id as id,
						users.id as student_id,
						curriculums.name as name,
						curriculums.id as curriculum_id,
						users.name as student_name,
						groups.id as group_id,
						groups.name as group_name
							`).
		Joins("inner join groups on groups.id=selective_courses.group_id ").
		Joins("inner join curriculums on curriculums.id=selective_course.curriculum_id").
		Joins("inner join users on selective_courses.student_id=users.id and users.is_student=?", true).
		Where("curriculums.id=?", id).Find(list)
	s.log.Debug(list)
	if res.Error != nil {
		s.log.Error(res.Error)
		return nil, res.Error
	}
	return list, nil
}
