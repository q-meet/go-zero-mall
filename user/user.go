package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero/mall/user/internal/config"
	"go-zero/mall/user/internal/server"
	"go-zero/mall/user/internal/svc"
	"go-zero/mall/user/types/user"
	"rpc-common/log/zapx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	writer, err := zapx.InitLogger()
	logx.Must(err)
	logx.SetWriter(writer)

	//rpc log,grpc的全局拦截器
	s.AddUnaryInterceptors(svc.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
