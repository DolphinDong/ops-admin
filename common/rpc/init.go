package rpc

import (
	"github.com/DolphinDong/ops-admin/common/rpc/clients/adminclient"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	AdminClient admin.AdminClient
)

func MustInitRpcClient(adminConfig *zrpc.RpcClientConf) {
	if adminConfig != nil {
		client, err := zrpc.NewClient(*adminConfig)
		if err != nil {
			logger.ZapLogger.Fatalf("New admin client failed: %+v", errors.WithStack(err))
		}
		AdminClient = adminclient.NewAdmin(client)
	}
	return
}
