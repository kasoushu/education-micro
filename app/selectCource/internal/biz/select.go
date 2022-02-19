package biz

import (
	"education/app/selectCource/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type SelectRepo interface {
}

type SelectCase struct {
	selectRepo SelectRepo
	log        *log.Helper
}

func NewSelectCase(repo SelectRepo, logger log.Logger) service.SelectCase {

}
