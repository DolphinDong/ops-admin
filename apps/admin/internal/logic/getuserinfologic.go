package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *pb.GetUserInfoReq) (*pb.GetUserInfoRes, error) {
	var result = &pb.GetUserInfoRes{}
	db := l.svcCtx.GetDB(l.ctx)
	userId := l.svcCtx.GetUserId(l.ctx)
	logger := l.svcCtx.Logger(l.ctx)
	var user *admin.User
	err := db.Model(&admin.User{}).Preload("Roles").Where("id=? and status = ? ", userId, admin.UserStatusEnable).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("获取用户信息失败，用户未找到")
			return nil, errors.New("获取用户信息失败，用户未找到")
		}
		logger.Errorf("查询用户信息失败：%v", errors.WithStack(err))
		return nil, errors.Errorf("查询用户信息失败：%s", err.Error())
	}
	result.Id = int64(user.ID)
	result.Username = user.Username
	result.NickName = user.NickName
	result.Email = user.Email
	result.Tel = user.Tel
	for _, r := range user.Roles {
		result.Roles = append(result.Roles, r.Identify)
	}
	logger.Error("查询用户信息成功")
	return result, nil
}
