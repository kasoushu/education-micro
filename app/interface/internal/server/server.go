package server

import (
	"education/app/interface/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewConsulRegister)

func NewConsulRegister(conf *conf.AppConfig) registry.Registrar {
	apiConfig := consulApi.DefaultConfig()
	if conf.Consul.Address != "" {
		apiConfig.Address = conf.Consul.Address
	}

	if conf.Consul.Scheme != "" {
		apiConfig.Scheme = conf.Consul.Scheme
	}
	apiCli, err := consulApi.NewClient(apiConfig)
	if err != nil {
		panic(err)
	}
	rs := consul.New(apiCli, consul.WithHealthCheck(false))
	return rs
}
