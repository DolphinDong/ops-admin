package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/pkg/errors"
)

type AddOrUpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateApiLogic {
	return &AddOrUpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateApiLogic) AddOrUpdateApi(req *pb.AddOrUpdateApiReq) (*pb.AddOrUpdateApiRes, error) {
	db := l.svcCtx.GetDB(l.ctx)
	logger := l.svcCtx.Logger(l.ctx)
	for _, inApi := range req.Apis {
		api := &admin.Api{
			Path:    inApi.Path,
			Method:  inApi.Method,
			Handler: inApi.Handler,
		}
		var count int64
		var a *admin.Api
		err := db.Model(&admin.Api{}).Where("path = ? and  method = ?", inApi.Path, inApi.Method).Count(&count).Find(&a).Error
		if err != nil {
			logger.Errorf("查询Api信息失败：%+v", errors.WithStack(err))
		}
		if count == 0 {
			err = db.Model(&admin.Api{}).Create(&api).Error
		} else {
			err = db.Model(&admin.Api{}).Where("id=?", a.Id).Updates(&api).Error
		}
		if err != nil {
			logger.Errorf("添加或者更改Api信息失败：%+v", errors.WithStack(err))
		}
	}
	return &pb.AddOrUpdateApiRes{}, nil
}
