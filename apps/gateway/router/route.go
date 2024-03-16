package router

import (
	"github.com/DolphinDong/ops-admin/apps/gateway/internal/service/admin"
	"github.com/gin-gonic/gin"
)

const (
	ApiV1Prefix = "/api/v1"
)

func InitRoute(engin *gin.Engine) {
	v1 := engin.Group(ApiV1Prefix)
	initAdminRoute(v1)
}

func initAdminRoute(route *gin.RouterGroup) {
	adminApi := admin.AdminApi{}
	adminRoute := route.Group("admin")
	{
		adminRoute.POST("login", adminApi.UserLogin)
	}
}
