package logic

import (
	"context"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *admin.LogoutReq) (*admin.LogoutRes, error) {
	// todo: add your logic here and delete this line

	return &admin.LogoutRes{}, nil
}
