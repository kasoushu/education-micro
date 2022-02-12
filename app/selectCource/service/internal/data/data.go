package data

import (
	"context"
	"education/app/selectCource/service/internal/conf"
	"education/app/selectCource/service/internal/data/ent"
	"education/app/selectCource/service/internal/data/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	log := log.NewHelper(log.With(logger, "module", "data"))
	log.Info("ent Client init")

	cli, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	data := Data{db: cli}
	cleanup := func() {
		log.Info("closing the data resources ")
		if err := data.db.Close(); err != nil {
			log.Error("close data db error")
		}
	}

	_, err = data.db.User.Create().SetName("liming").SetPassword("0000").SetIsAdmin(true).SetIsStudent(false).SetIsTeacher(true).Save(context.Background())
	if err != nil {
		log.Error(err)
	}
	return &Data{}, cleanup, nil
}
