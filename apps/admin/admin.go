package main

import (
	"flag"
	"fmt"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/config"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/server"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the admin config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterAdminServer(grpcServer, server.NewAdminServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
