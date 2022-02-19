package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Id            uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	Name          string
	MajorId       uint64
	DepartmentId  uint64
	StudentNumber int
}
