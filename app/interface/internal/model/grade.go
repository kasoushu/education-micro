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

type CreateGrade struct {
	StudentId    uint64
	CurriculumId uint64
	Score        int32
	StudyLevel   string
	GradePoint   int32
	Note         string
	GradeYear    int32
	Term         int32
}

type SingleGrade struct {
	Id             uint64
	CurriculumId   uint64
	CurriculumName string
	Credit         int32
	Category       string
	ExamWay        string
	StudyLevel     string
	Score          int32
	GainCredit     int32
	GradePoint     int32
	Note           string
	StudentName    string
	StudentId      uint64
}
type GradeQueryByCurriculumOnOneTerm struct {
	CurriculumId uint64
	StudentId    uint64
	GradeYear    int32
	Term         int32
}

type GradeQueryByStudentIdOnOneTerm struct {
	StudentId uint64
	GradeYear int32
	Term      int32
}
type GradeQueryByGroupIdOnOneTerm struct {
	GroupId      uint64
	CurriculumId uint64
	GradeYear    int32
	Term         int32
}
