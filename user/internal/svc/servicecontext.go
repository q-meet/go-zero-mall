package svc

import (
	"go-zero/mall/user/dao"
	"go-zero/mall/user/dao/database"
	"go-zero/mall/user/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo *dao.UserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(database.Connect(c.Mysql.DataSource, c.CacheRedis)),
	}
}
