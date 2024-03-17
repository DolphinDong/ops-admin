package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/DolphinDong/ops-admin/pkg/tools"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type CheckTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckTokenLogic) CheckToken(req *pb.CheckTokenReq) (*pb.CheckTokenRes, error) {
	db := l.svcCtx.GetDB(l.ctx)
	var user *admin.User
	claims, err := tools.ParseToken(req.Token)
	if err != nil {
		// token 过期
		if strings.Contains(err.Error(), "token is expired by") {
			l.svcCtx.Logger(l.ctx).Infof("Token已过期，请重新登陆")
			return &pb.CheckTokenRes{Success: false, Message: "Token已过期，请重新登陆"}, nil
		}
		l.svcCtx.Logger(l.ctx).Errorf("Parse token failed: %+v", errors.WithStack(err))
		return nil, errors.Errorf("Parse token error: %s", err.Error())
	}

	err = db.Model(&admin.User{}).Where("id=?", claims.Issuer).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.svcCtx.Logger(l.ctx).Warn("用户不存在")
			return &pb.CheckTokenRes{Success: false, Message: "用户不存在"}, nil
		}
		l.svcCtx.Logger(l.ctx).Errorf("Query user info failed: %+v", errors.WithStack(err))
		return nil, errors.Errorf("Query user info failed：%s", err)
	}

	if user.Status == admin.UserStatusDisable {
		l.svcCtx.Logger(l.ctx).Warnf("用户id=%v已被禁用，请联系管理员", claims.Issuer)
		return &pb.CheckTokenRes{Success: false, Message: "用户已被禁用，请联系管理员"}, nil
	}

	userId, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		l.svcCtx.Logger(l.ctx).Errorf("Convert userId failed: %+v", errors.WithStack(err))
		return nil, errors.Errorf("Convert userId error: %s", err.Error())
	}
	return &pb.CheckTokenRes{
		Success: true,
		UserId:  int64(userId),
	}, nil
}
