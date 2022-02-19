package model

import (
	"gorm.io/gorm"
	"time"
)

type Curriculum struct {
	gorm.Model
	Name        string
	Id          uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	ClassroomId uint64
	StartTime   time.Time
	EndTime     time.Time
	Credit      int
	SchoolHour  int
	Category    string
	Status      int    // 0,1,2,3,4
	ExamWay     string //
	TeacherId   uint64
}

type CurriculumItem struct {
}
