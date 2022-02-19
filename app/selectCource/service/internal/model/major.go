package model

import "gorm.io/gorm"

type Major struct {
	gorm.Model
	Id           uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	Name         string
	DepartmentId uint64
	Description  string
}
