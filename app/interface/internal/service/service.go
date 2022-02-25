package service

import (
	iv1 "education/api/v1/interface"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is interface providerso.

var ProviderSet = wire.NewSet(NewService)

type InterfaceService struct {
	iv1.UnimplementedEducationInterfaceServer
	userCase   UserCase
	courseCase CourseCase
	selectCase SelectCase
	gradeCase  GradeCase
	log        *log.Helper
}

func NewService(u UserCase, c CourseCase, s SelectCase, g GradeCase, l log.Logger) *InterfaceService {
	return &InterfaceService{
		userCase:   u,
		courseCase: c,
		selectCase: s,
		gradeCase:  g,
		log:        log.NewHelper(log.With(l, "module", "interface-interface")),
	}
}
