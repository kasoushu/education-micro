package data

import (
	"education/app/selectCource/service/internal/conf"
	"education/app/selectCource/service/internal/model"
	"github.com/go-kratos/kratos/v2/log"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDb)

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
	db = db.Set("gorm:table_options", "AUTOINCREMENT=100000000")
	err = db.AutoMigrate(&model.Class{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Group{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Grade{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Department{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Major{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Curriculum{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.SelectiveCourse{})
	if err != nil {
		log.Fatal(err)
	}
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
