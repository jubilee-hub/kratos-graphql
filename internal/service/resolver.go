package service

import (
	"arkgo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log    *log.Helper
	userUc *biz.UserUsecase
}

func NewResolverService(
	logger log.Logger,
	userUc *biz.UserUsecase,
) *Resolver {
	return &Resolver{
		log:    log.NewHelper(logger),
		userUc: userUc,
	}
}
