package config

import "github.com/zeromicro/go-zero/zrpc"

type Mysql struct {
	HostName    string
	Port        int
	Username    string
	Password    string
	Database    string
	MaxIdle     int    `json:"MaxIdle,default=10"`
	MaxOpen     int    `json:"MaxOpen,default=100"`
	SqlLogLevel string `json:"SqlLogLevel,default=warn,options=info|warn|error|silent"`
}
type ServerConfig struct {
	Admin *zrpc.RpcClientConf
}
