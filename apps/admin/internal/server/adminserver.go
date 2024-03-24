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

func (s *AdminServer) Logout(ctx context.Context, in *admin.LogoutReq) (*admin.LogoutRes, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *AdminServer) CheckToken(ctx context.Context, in *admin.CheckTokenReq) (*admin.CheckTokenRes, error) {
	l := logic.NewCheckTokenLogic(ctx, s.svcCtx)
	return l.CheckToken(in)
}

func (s *AdminServer) CheckPermission(ctx context.Context, in *admin.CheckPermissionReq) (*admin.CheckPermissionRes, error) {
	l := logic.NewCheckPermissionLogic(ctx, s.svcCtx)
	return l.CheckPermission(in)
}

// 列出当前 用户拥有的菜单
func (s *AdminServer) GetMenuList(ctx context.Context, in *admin.GetMenuListReq) (*admin.GetMenuListRes, error) {
	l := logic.NewGetMenuListLogic(ctx, s.svcCtx)
	return l.GetMenuList(in)
}

func (s *AdminServer) GetUserInfo(ctx context.Context, in *admin.GetUserInfoReq) (*admin.GetUserInfoRes, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *AdminServer) GetPermCode(ctx context.Context, in *admin.GetPermCodeReq) (*admin.GetPermCodeRes, error) {
	l := logic.NewGetPermCodeLogic(ctx, s.svcCtx)
	return l.GetPermCode(in)
}

func (s *AdminServer) AddOrUpdateApi(ctx context.Context, in *admin.AddOrUpdateApiReq) (*admin.AddOrUpdateApiRes, error) {
	l := logic.NewAddOrUpdateApiLogic(ctx, s.svcCtx)
	return l.AddOrUpdateApi(in)
}
