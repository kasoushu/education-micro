package data

import (
	"context"
	"education/app/selectCource/internal/biz"
	"education/app/selectCource/internal/model"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type CourseRepo struct {
	data *Data
	log  *log.Helper
}

func NewCourseRepo(d *Data, logger log.Logger) biz.CourseRepo {
	return &CourseRepo{
		data: d,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (c *CourseRepo) Create(ctx context.Context, curriculum *model.Curriculum) error {
	cc := model.Curriculum{}
	res := c.data.db.WithContext(ctx).First(cc, "name=?", curriculum.Name)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.CURRICULUM_HAD_EXISTED
	}
	if res = c.data.db.WithContext(ctx).Create(curriculum); res.Error != nil {
		c.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (c *CourseRepo) Update(ctx context.Context, curriculum *model.Curriculum, id uint64) error {
	cc := model.Curriculum{}
	res := c.data.db.WithContext(ctx).First(cc, id)
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

func (c *CourseRepo) Delete(ctx context.Context, id uint64) error {
	res := c.data.db.WithContext(ctx).Delete(&model.Curriculum{}, id)
	if res.Error != nil {
		c.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (c *CourseRepo) GetSingle(ctx context.Context, id uint64) (*model.CurriculumItem, error) {
	cc := model.CurriculumItem{}
	res := c.data.db.Table("curriculums").
		Select(`curriculums.id as id,
						classrooms.id as classroom_id,
						users.id as teacher_id,
						classrooms.name as classroom_name,
						curriculums.name as name,
						users.name as teacher_name,
						curriculums.grade_year as grade_year,
						curriculums.term as term,
						curriculums.category as category,
						curriculums.exam_way as exam_way,
						curriculums.school_hour school_hour,
						curriculums.status as status,
						curriculums.credit as credit,
							`).
		Joins("inner join classrooms on curriculums.classroom_id = classroom.id ").
		Joins("inner join users on curriculums.teacher_id=users.id and users.is_teacher=?", true).
		Where("curriculums.id=?", id).First(&cc)
	c.log.Debug(cc)
	if res.Error != nil {
		c.log.Error(res.Error)
		return nil, res.Error
	}
	return &cc, nil
}

func (c *CourseRepo) GetListByTeacherId(ctx context.Context, id uint64) ([]*model.CurriculumItem, error) {
	var list []*model.CurriculumItem
	result := c.data.db.Table("curriculums").
		Select(`curriculums.id as id,
						classrooms.id as classroom_id,
						users.id as teacher_id,
						classrooms.name as classroom_name,
						curriculums.name as name,
						users.name as teacher_name,
						curriculums.grade_year as grade_year,
						curriculums.term as term,
						curriculums.category as category,
						curriculums.exam_way as exam_way,
						curriculums.school_hour school_hour,
						curriculums.status as status,
						curriculums.credit as credit,
							`).
		Joins("inner join classrooms on curriculums.classroom_id = classroom.id ").
		Joins("inner join users on curriculums.teacher_id=users.id and users.is_teacher=?", true).
		Where("users.id=?", id).Find(list)
	if result.Error != nil {
		return nil, result.Error
	}
	c.log.Debug(list)
	return list, nil
}
