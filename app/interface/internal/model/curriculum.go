package model

import (
	"gorm.io/gorm"
)

type Curriculum struct {
	gorm.Model
	Name        string
	Id          uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	ClassroomId uint64
	GradeYear   int32
	Term        int32
	Credit      int32
	SchoolHour  int32
	Category    string
	Status      int32  // 0未选,1 可选 ,2 未开课 ,3 未结课 ,4已结课
	ExamWay     string //
	TeacherId   uint64
}

type CurriculumItem struct {
	Id            uint64
	ClassroomId   uint64
	ClassroomName string
	TeacherId     uint64
	TeacherName   uint64
	Name          string
	GradeYear     int32
	Term          int32
	Credit        int32
	SchoolHour    int32
	Category      string
	Status        int32
	ExamWay       string
}
