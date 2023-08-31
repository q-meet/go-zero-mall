package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc-common/log/interceptor"
	"rpc-common/log/zapx"

	"go-zero/mall/userscore/internal/config"
	"go-zero/mall/userscore/internal/server"
	"go-zero/mall/userscore/internal/svc"
	"go-zero/mall/userscore/types/userscore"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userscore.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userscore.RegisterUserscoreServer(grpcServer, server.NewUserscoreServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	writer, err := zapx.InitLogger()
	//writer, err := zapx.NewCore()
	logx.Must(err)
	logx.SetWriter(writer)

	//rpc log,grpc的全局拦截器
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
