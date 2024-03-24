package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/common/consts"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	"github.com/DolphinDong/ops-admin/pkg/tools"
	"github.com/pkg/errors"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
)

type GetPermCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermCodeLogic {
	return &GetPermCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermCodeLogic) GetPermCode(req *pb.GetPermCodeReq) (*pb.GetPermCodeRes, error) {
	var result = &pb.GetPermCodeRes{}
	db := l.svcCtx.GetDB(l.ctx)
	userId := l.svcCtx.GetUserId(l.ctx)
	logger := l.svcCtx.Logger(l.ctx)
	var user *admin.User
	logger.Infof("查询用户（%v）权限编码", userId)
	err := db.Model(&admin.User{}).Preload("Roles", "status=?", admin.RoleStatusEnable).
		Preload("Roles.Menus", "status=?", admin.MenuStatusEnable).
		Where("id=?", userId).Find(&user).Error
	if err != nil {
		logger.Errorf("List user menu failed: %+v", errors.WithStack(err))
		return nil, errors.WithMessage(err, "List user menu failed")
	}
	for _, role := range user.Roles {
		// 如果当前用户未管理员的话就直接返回*，代表所有的权限
		if role.Identify == consts.AdminRoleIdentify {
			result.PermissionCodes = []string{"*"}
			return result, nil
		}
		for _, menu := range role.Menus {
			if menu.Permission != "" {
				result.PermissionCodes = append(result.PermissionCodes, menu.Permission)
			}
		}
	}
	result.PermissionCodes = tools.RemoveDuplication(result.PermissionCodes)
	return result, nil
}
