package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero/mall/user/Api/internal/config"
	"go-zero/mall/user/Api/internal/handler"
	"go-zero/mall/user/Api/internal/svc"
	"rpc-common/log/zapx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	writer, err := zapx.InitLogger()
	logx.Must(err)
	logx.SetWriter(writer)

	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.Infow(fmt.Sprintf("Starting server at %s:%d...\n", c.Host, c.Port), logx.Field("host", c.Host))
	server.Start()
}
