package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero/mall/user/Api/internal/config"
	"go-zero/mall/user/Api/internal/middleware"
	"go-zero/mall/user/userclient"
	"go-zero/mall/userscore/userscoreclient"
)

type ServiceContext struct {
	Config            config.Config
	UserRpc           userclient.User
	UserScoreRpc      userscoreclient.Userscore
	ExampleMiddleware *middleware.ExampleMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserRpc:           userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserScoreRpc:      userscoreclient.NewUserscore(zrpc.MustNewClient(c.UserScoreRpc)),
		ExampleMiddleware: middleware.NewExampleMiddleware(),
	}
}
