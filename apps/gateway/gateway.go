package main

import (
	"flag"
	"github.com/DolphinDong/ops-admin/apps/gateway/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the gateway config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
}
