package model

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Id           uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	StudentId    uint64
	CurriculumId uint64
	Score        int32
	StudyLevel   string
	GradePoint   int32
	Credit       int32
	Note         string
	GradeYear    int32
	Term         int32
}
