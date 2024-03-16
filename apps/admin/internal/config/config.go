package config

import (
	"github.com/DolphinDong/ops-admin/common/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql config.Mysql
}
