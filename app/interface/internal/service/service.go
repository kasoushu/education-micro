package service

import "github.com/google/wire"

// ProviderSet is interface providerso.

var ProviderSet = wire.NewSet(NewUserService)
