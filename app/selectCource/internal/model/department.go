package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Id          uint64 `gorm:"<-:false;primaryKey;autoIncrement"`
	Name        string
	MajorNumber int
	Description string
}
