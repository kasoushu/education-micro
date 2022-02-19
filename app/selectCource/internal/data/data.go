package data

import (
	"education/app/selectCource/internal/conf"
	model2 "education/app/selectCource/internal/model"
	"github.com/go-kratos/kratos/v2/log"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDb, NewCourseRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewGormDb(conf *conf.AppConfig, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Data.Database.Source), &gorm.Config{})
	l := log.NewHelper(logger)
	if err != nil {
		l.Fatal(err)
	}
	log := log.NewHelper(logger)
	log.Info("db initiating!")
	db = db.Set("gorm:table_options", "AUTO_INCREMENT=100000000")
	err = db.AutoMigrate(&model2.Class{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model2.Group{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model2.Grade{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model2.Department{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model2.Major{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model2.Curriculum{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model2.SelectiveCourse{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info("db initiated")
	return db
}

// NewData .
func NewData(c *conf.AppConfig, db *gorm.DB, logger log.Logger) (*Data, func(), error) {

	return &Data{
			db:  db,
			log: log.NewHelper(log.With(logger, "module", "data")),
		}, func() {
			log.NewHelper(logger).Info("cleaning!\n closing!")
		}, nil
}
