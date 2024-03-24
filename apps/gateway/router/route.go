package router

import (
	"github.com/DolphinDong/ops-admin/apps/gateway/internal/service/admin"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	ApiV1Prefix = "/api/v1"
)

func InitRoute(engin *gin.Engine) {
	v1 := engin.Group(ApiV1Prefix)
	initAdminRoute(v1)
	err := AddOrUpdateApis(engin)
	if err != nil {
		logger.ZapLogger.Errorf("添加或者修改API信息失败：%+v", errors.WithStack(err))
	}
	logger.ZapLogger.Info("添加或者修改API信息成功")
}

func initAdminRoute(route *gin.RouterGroup) {
	adminApi := admin.AdminApi{}
	adminRoute := route.Group("admin")
	{
		adminRoute.POST("login", adminApi.UserLogin)
		adminRoute.GET("logout", adminApi.UserLogout)
		adminRoute.GET("getPermCode", adminApi.GetPermCode)
		adminRoute.GET("getUserInfo", adminApi.GetUserInfo)
		adminRoute.GET("getMenuList", adminApi.GetMenuList)
		adminRoute.POST("ping/:id", adminApi.Ping)
	}
}
