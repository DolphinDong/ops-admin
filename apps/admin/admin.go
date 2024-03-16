package main

import (
	"flag"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/config"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/server"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the admin config file")

var (
	Logger *zap.SugaredLogger
)

func init() {
	logger.SetupZap()
	Logger = logger.ZapLogger
}

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
	go func() {
		Logger.Infof("Starting rpc server at %s...", c.ListenOn)
		s.Start()
	}()
	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	Logger.Info("Start Shutdown Server")
	s.Stop()
	Logger.Info("Server Shutdown  success")
}
