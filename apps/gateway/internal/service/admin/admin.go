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
		a.OK(nil, res.Message)
	} else {
		a.OK(res, "登录成功")
	}
}

func (a AdminApi) Ping(c *gin.Context) {
	api.OK(c, nil, "pong")
}
