//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"education/app/selectCource/internal/biz"
	"education/app/selectCource/internal/conf"
	"education/app/selectCource/internal/data"
	"education/app/selectCource/internal/server"
	"education/app/selectCource/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.AppConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
	//panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
