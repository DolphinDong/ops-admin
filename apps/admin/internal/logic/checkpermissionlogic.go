package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/casbin"
	"github.com/pkg/errors"
	"strconv"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
)

type CheckPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPermissionLogic {
	return &CheckPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckPermissionLogic) CheckPermission(req *pb.CheckPermissionReq) (*pb.CheckPermissionRes, error) {
	// TODO 判断当前用户是否为管理员，如果是管理员的话则直接返回校验通过

	enforce, err := casbin.GetEnforcer().Enforce(strconv.Itoa(int(req.UserId)), req.Url, req.Method)
	if err != nil {
		l.svcCtx.Logger(l.ctx).Errorf("Check permission error: %v", err)
		return nil, errors.Errorf("Check permission error: %s", err)
	}
	return &pb.CheckPermissionRes{
		Success: enforce,
	}, nil
}
