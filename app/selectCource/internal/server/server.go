package server

import (
	"education/app/selectCource/internal/conf"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/wire"
	//"education/app/selectCource/interface/internal/conf"
	"github.com/go-kratos/kratos/v2/registry"
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
	apicli, err := consulApi.NewClient(apiConfig)
	if err != nil {
		panic(err)
	}
	rs := consul.New(apicli, consul.WithHealthCheck(false))
	return rs
}
