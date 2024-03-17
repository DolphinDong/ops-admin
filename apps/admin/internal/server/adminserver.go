// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package server

import (
	"context"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/logic"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminServer) Login(ctx context.Context, in *admin.LoginReq) (*admin.LoginRes, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *AdminServer) CheckToken(ctx context.Context, in *admin.CheckTokenReq) (*admin.CheckTokenRes, error) {
	l := logic.NewCheckTokenLogic(ctx, s.svcCtx)
	return l.CheckToken(in)
}

func (s *AdminServer) CheckPermission(ctx context.Context, in *admin.CheckPermissionReq) (*admin.CheckPermissionRes, error) {
	l := logic.NewCheckPermissionLogic(ctx, s.svcCtx)
	return l.CheckPermission(in)
}
