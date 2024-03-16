package config

import "github.com/zeromicro/go-zero/zrpc"

type Mysql struct {
	HostName string
	Port     int64
	Username string
	Password string
	Database string
}
type ServerConfig struct {
	Admin *zrpc.RpcClientConf
}
