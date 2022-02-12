package server

import (
	"education/app/selectCource/service/internal/conf"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/wire"
	//"education/app/selectCource/service/internal/conf"
	"github.com/go-kratos/kratos/v2/registry"
	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewConsulRegister)

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
