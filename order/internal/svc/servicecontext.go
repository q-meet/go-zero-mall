package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero/mall/order/internal/config"
	"go-zero/mall/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
