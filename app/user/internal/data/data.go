package data

import (
	"context"
	"education/app/user/internal/conf"
	"education/app/user/internal/data/ent"
	"education/app/user/internal/data/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	myLog := log.NewHelper(log.With(logger, "module", "data"))
	cli, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		myLog.Fatal(err)
	}
	myLog.Info("ent Client init successful!")
	if err := cli.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		myLog.Fatalf("failed creating schema resources: %v", err)
	}
	data := Data{db: cli}
	cleanup := func() {
		myLog.Info("closing the data resources ")
		if err := data.db.Close(); err != nil {
			myLog.Error("close data db error")
		}
	}

	//_, err = data.db.User.Create().SetName("liming").SetPassword("0000").SetIsAdmin(true).SetIsStudent(false).SetIsTeacher(true).Save(context.Background())
	//if err != nil {
	//	myLog.Error(err)
	//}
	return &data, cleanup, nil
}
