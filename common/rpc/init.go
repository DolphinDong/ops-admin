package rpc

import (
	"github.com/DolphinDong/ops-admin/common/rpc/clients/adminclient"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
)

var (
	AdminClient admin.AdminClient
)

func MustInitRpcClient(adminConfig *zrpc.RpcClientConf) {
	if adminConfig != nil {
		client, err := zrpc.NewClient(*adminConfig)
		if err != nil {
			logx.Errorf("New admin client failed: %+v", errors.WithStack(err))
			os.Exit(1)
		}
		AdminClient = adminclient.NewAdmin(client)
	}
	return
}
