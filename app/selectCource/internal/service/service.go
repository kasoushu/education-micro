package service

import "github.com/google/wire"

// ProviderSet is interface providers.
var ProviderSet = wire.NewSet(NewCourseService)
