package admin

import (
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/DolphinDong/ops-admin/common/rpc"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type AdminApi struct {
	api.Api
}

func (a AdminApi) UserLogin(c *gin.Context) {
	var req admin.LoginReq
	err := a.MakeContext(c).Bind(&req, binding.JSON, binding.Query).Validate(&req).Errors
	if err != nil {
		a.Error(http.StatusBadRequest, err, err.Error())
		a.Logger.Error(err.Error())
		return
	}
	res, err := rpc.AdminClient.Login(a.MetadataContext, &req)
	if err != nil {
		a.Logger.Errorf("Call admin server login api failed: %+v", err)
		api.Error(c, http.StatusInternalServerError, err, err.Error())
		return
	}
	if !res.Success {
		a.Error(http.StatusBadRequest, nil, res.Message)
	} else {
		a.OK(res, "登录成功")
	}
}

func (a AdminApi) Ping(c *gin.Context) {
	api.OK(c, nil, "pong")
}

func (a AdminApi) GetMenuList(c *gin.Context) {
	var req admin.GetMenuListReq
	err := a.MakeContext(c).Bind(&req, binding.JSON, binding.Query).Validate(&req).Errors
	if err != nil {
		a.Error(http.StatusBadRequest, err, err.Error())
		a.Logger.Error(err.Error())
		return
	}
	res, err := rpc.AdminClient.GetMenuList(a.MetadataContext, &req)
	if err != nil {
		a.Logger.Errorf("Call admin server GetMenuList failed: %+v", err)
		api.Error(c, http.StatusInternalServerError, err, err.Error())
		return
	}
	a.OK(res.Menu, "查询菜单成功")
}

func (a AdminApi) GetUserInfo(c *gin.Context) {
	var req admin.GetUserInfoReq
	err := a.MakeContext(c).Bind(&req, binding.JSON, binding.Query).Validate(&req).Errors
	if err != nil {
		a.Error(http.StatusBadRequest, err, err.Error())
		a.Logger.Error(err.Error())
		return
	}
	res, err := rpc.AdminClient.GetUserInfo(a.MetadataContext, &req)
	if err != nil {
		a.Logger.Errorf("Call admin server GetUserInfo failed: %+v", err)
		api.Error(c, http.StatusInternalServerError, err, err.Error())
		return
	}
	a.OK(res, "查询用户信息成功")
}

func (a AdminApi) GetPermCode(c *gin.Context) {
	var req admin.GetPermCodeReq
	err := a.MakeContext(c).Bind(&req, binding.JSON, binding.Query).Validate(&req).Errors
	if err != nil {
		a.Error(http.StatusBadRequest, err, err.Error())
		a.Logger.Error(err.Error())
		return
	}
	res, err := rpc.AdminClient.GetPermCode(a.MetadataContext, &req)
	if err != nil {
		a.Logger.Errorf("Call admin server GetPermCode failed: %+v", err)
		api.Error(c, http.StatusInternalServerError, err, err.Error())
		return
	}
	a.OK(res.PermissionCodes, "查询用户权限码成功")
}

func (a AdminApi) UserLogout(c *gin.Context) {
	var req admin.LogoutReq
	err := a.MakeContext(c).Bind(&req, binding.JSON, binding.Query).Validate(&req).Errors
	if err != nil {
		a.Error(http.StatusBadRequest, err, err.Error())
		a.Logger.Error(err.Error())
		return
	}
	_, err = rpc.AdminClient.Logout(a.MetadataContext, &req)
	if err != nil {
		a.Logger.Errorf("Call admin server Logout failed: %+v", err)
		api.Error(c, http.StatusInternalServerError, err, err.Error())
		return
	}
	a.OK(nil, "退出登录成功")
}
