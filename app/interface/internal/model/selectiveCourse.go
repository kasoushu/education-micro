package model

import "gorm.io/gorm"

type SelectiveCourse struct {
	gorm.Model
	Id           uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	CurriculumId uint64
	GroupId      uint64
	StudentId    uint64
}

type SelectInfo struct {
	Id             uint64
	CurriculumId   uint64
	GroupId        uint64
	StudentId      uint64
	StudentName    string
	GroupName      string
	CurriculumName string
}
