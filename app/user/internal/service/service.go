package service

import "github.com/google/wire"

// ProviderSet is service providerso.

var ProviderSet = wire.NewSet(NewUserService)
