package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/DolphinDong/ops-admin/pkg/tools"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	zap.SugaredLogger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *pb.LoginReq) (*pb.LoginRes, error) {
	var res = &pb.LoginRes{}
	db := l.svcCtx.GetDB(l.ctx)
	var user *admin.User
	err := db.Model(&admin.User{}).Where("username=? and password=?", req.Username, tools.GetEncryptedPassword(req.Password)).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.Success = false
			res.Message = "用户名或密码不正确"
			l.svcCtx.Logger(l.ctx).Warn("登录失败，用户名或者密码不正确")
			return res, nil
		} else {
			l.svcCtx.Logger(l.ctx).Errorf("Login failed: %+v", errors.WithStack(err))
			return nil, errors.Errorf("login failed：%s", err)
		}
	}

	if user.Status == admin.UserStatusDisable {
		res.Success = false
		res.Message = "用户已被禁用，请联系管理员"
		l.svcCtx.Logger(l.ctx).Warnf("用户 %v 已被禁用，请联系管理员", req.Username)
		return res, nil
	}
	token, err := tools.CreateToken(strconv.Itoa(int(user.ID)), l.svcCtx.Config.TokenExpireSec)
	if err != nil {
		l.svcCtx.Logger(l.ctx).Errorf("Create token failed: %+v", errors.WithStack(err))
		return nil, errors.Errorf("Token创建失败：%s", err)
	}
	return &pb.LoginRes{
		Success: true,
		Token:   token,
	}, nil
}
