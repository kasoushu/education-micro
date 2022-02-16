package server

import (
	"education/app/user/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewConsulRegister)

func NewConsulRegister(conf *conf.Consul) registry.Registrar {
	apiConfig := consulApi.DefaultConfig()
	if conf.Address != "" {
		apiConfig.Address = conf.Address
	}

	if conf.Scheme != "" {
		apiConfig.Scheme = conf.Scheme
	}
	apicli, err := consulApi.NewClient(apiConfig)
	if err != nil {
		panic(err)
	}
	rs := consul.New(apicli, consul.WithHealthCheck(false))
	return rs
}
