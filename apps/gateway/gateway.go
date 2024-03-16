package main

import (
	"context"
	"flag"
	"github.com/DolphinDong/ops-admin/apps/gateway/internal/config"
	"github.com/DolphinDong/ops-admin/apps/gateway/middleware"
	"github.com/DolphinDong/ops-admin/apps/gateway/router"
	"github.com/DolphinDong/ops-admin/common/rpc"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the gateway config file")

var (
	Logger *zap.SugaredLogger
)

func init() {
	logger.SetupZap()
	Logger = logger.ZapLogger
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()
	var c config.Config
	Logger.Info("loading config file success")
	conf.MustLoad(*configFile, &c)
	rpc.MustInitRpcClient(c.ServerConfig.Admin)
	engine := gin.New()
	InitEngine(engine)
	//创建HTTP服务器
	server := &http.Server{
		Addr:    c.ListenOn,
		Handler: engine,
	}
	//启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Fatalf("Server start failed: %+v", errors.WithStack(err))
		}
	}()
	Logger.Infof("Server listen at : %v", c.ListenOn)
	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	Logger.Info("Start Shutdown Server")
	if err := server.Shutdown(ctx); err != nil {
		Logger.Fatalf("Server Shutdown failed: %+v", errors.WithStack(err))
	}
	Logger.Info("Server Shutdown  success")
}

func InitEngine(engin *gin.Engine) {
	Logger.Info("Init middleware success")
	middleware.InitMiddleware(engin)
	Logger.Info("Init route success")
	router.InitRoute(engin)
}
