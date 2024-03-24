package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/casbin"
	"github.com/DolphinDong/ops-admin/common/consts"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	userId := l.svcCtx.GetUserId(l.ctx)
	logger := l.svcCtx.Logger(l.ctx)
	db := l.svcCtx.GetDB(l.ctx)
	var user *admin.User
	err := db.Model(&admin.User{}).Preload("Roles", "identify = ?", consts.AdminRoleIdentify).Where("status =? and id = ?", admin.UserStatusEnable, userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在或者被禁用")
		}
		logger.Errorf("查询用户信息失败：%+v", errors.WithStack(err))
		return nil, errors.WithMessage(err, "查询用户信息失败")
	}
	if len(user.Roles) > 0 {
		return &pb.CheckPermissionRes{
			Success: true,
		}, nil
	}

	enforce, err := casbin.GetEnforcer().Enforce(strconv.Itoa(int(req.UserId)), req.Url, req.Method)
	if err != nil {
		l.svcCtx.Logger(l.ctx).Errorf("Check permission error: %v", err)
		return nil, errors.Errorf("Check permission error: %s", err)
	}
	return &pb.CheckPermissionRes{
		Success: enforce,
	}, nil
}
