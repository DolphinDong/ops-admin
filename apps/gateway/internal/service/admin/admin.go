package admin

import (
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type AdminApi struct {
	api.Api
}
type LoginReq struct {
	Username string `json:"username" form:"username" uri:"username" validate:"required"`
	Password string `json:"password" form:"password" uri:"password" validate:"required"`
}

func (a AdminApi) UserLogin(c *gin.Context) {
	var loginReq LoginReq
	err := a.MakeContext(c).Bind(&loginReq, binding.JSON, binding.Query).Validate(&loginReq).Errors
	if err != nil {
		a.Error(http.StatusBadRequest, err, err.Error())
		a.Logger.Error(err.Error())
		return
	}
	a.OK(nil, "login success")
}
