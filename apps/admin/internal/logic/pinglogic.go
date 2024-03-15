package logic

import (
	"context"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *admin.Request) (*admin.Response, error) {
	// todo: add your logic here and delete this line

	return &admin.Response{}, nil
}
