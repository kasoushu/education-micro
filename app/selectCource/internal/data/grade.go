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
		g.log.Error(res.Error)
		return res.Error
	}
	cc.Id = id
	res = g.data.db.WithContext(ctx).Model(cc).Updates(grade)
	if res.Error != nil {
		g.log.Error(res.Error)
		return res.Error
	}
	return nil
}

func (g *GradeRepo) GetGradeByCurriculum(ctx context.Context, GradeQuery *model.GradeQueryByCurriculumOnOneTerm) (*model.SingleGrade, error) {
	cc := model.SingleGrade{}
	res := g.data.db.Table("grades").
		Select(`
						grades.id as id
						curriculums.id as curriculum_id,
						curriculums.name as curriculum_name,
						curriculums.credit as credit,
						curriculums.category as category,
						curriculums.exam_way as exam_way,
						grades.study_level as study_level,
						score,
						grades.credit as gain_credit,
						grades.grade_point as grade_point
						grades.note as note,
						users.id as student_id,
						users.name as student_name,
							`).
		Joins("inner join curriculums on curriculums.id = grades.curriculum_id and curriculums.id=?", GradeQuery.CurriculumId).
		Joins("inner join users on curriculums.student_id=users.id and users.is_student=? and users.id=? ", true, GradeQuery.StudentId).
		Where("grade_year=? ,term=? ", GradeQuery.GradeYear, GradeQuery.Term).First(&cc)
	g.log.Debug(cc)
	if res.Error != nil {
		g.log.Error(res.Error)
		return nil, res.Error
	}
	return &cc, nil
}

func (g *GradeRepo) GetGradeByStudentId(ctx context.Context, GradeQuery *model.GradeQueryByStudentIdOnOneTerm) ([]*model.SingleGrade, error) {

	var list []*model.SingleGrade
	res := g.data.db.Table("grades").
		Select(`
						grades.id as id
						curriculums.id as curriculum_id,
						curriculums.name as curriculum_name,
						curriculums.credit as credit,
						curriculums.category as category,
						curriculums.exam_way as exam_way,
						grades.study_level as study_level,
						score,
						grades.credit as gain_credit,
						grades.grade_point as grade_point
						grades.note as note,
						users.id as student_id,
						users.name as student_name,
							`).
		Joins("inner join curriculums on curriculums.id = grades.curriculum_id").
		Joins("inner join users on curriculums.student_id=users.id and users.is_student=? and users.id=? ", true, GradeQuery.StudentId).
		Where("grade_year=? ,term=? ", GradeQuery.GradeYear, GradeQuery.Term).Find(list)
	g.log.Debug(list)
	if res.Error != nil {
		g.log.Error(res.Error)
		return nil, res.Error
	}
	return list, nil
}

func (g *GradeRepo) GetGradeByGroupId(ctx context.Context, GradeQuery *model.GradeQueryByGroupIdOnOneTerm) ([]*model.SingleGrade, error) {
	var list []*model.SingleGrade
	res := g.data.db.Table("grades").
		Select(`
						grades.id as id
						curriculums.id as curriculum_id,
						curriculums.name as curriculum_name,
						curriculums.credit as credit,
						curriculums.category as category,
						curriculums.exam_way as exam_way,
						grades.study_level as study_level,
						score,
						grades.credit as gain_credit,
						grades.grade_point as grade_point
						grades.note as note,
						users.id as student_id,
						users.name as student_name,
							`).
		Joins("inner join curriculums on curriculums.id = grades.curriculum_id and curriculum.id=?", GradeQuery.CurriculumId).
		Joins("inner join users on curriculums.student_id=users.id and users.is_student=?", true).
		Where("grade_year=? ,term=? ", GradeQuery.GradeYear, GradeQuery.Term).
		Where("users.id in (?)", g.data.db.Table("selective_courses").Select("student_id").Where("selective_courses.group_id=?", GradeQuery.GroupId)).
		Find(list)
	g.log.Debug(list)
	if res.Error != nil {
		g.log.Error(res.Error)
		return nil, res.Error
	}
	return list, nil
}
