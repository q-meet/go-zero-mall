package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero/mall/user/Api/internal/config"
	"go-zero/mall/user/Api/internal/middleware"
	"go-zero/mall/user/userclient"
)

type ServiceContext struct {
	Config            config.Config
	UserRpc           userclient.User
	ExampleMiddleware *middleware.ExampleMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserRpc:           userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ExampleMiddleware: middleware.NewExampleMiddleware(),
	}
}
