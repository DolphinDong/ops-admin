package middleware

import (
	"github.com/DolphinDong/ops-admin/apps/gateway/router"
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	TokenKeyInHeader    = "Authorization"
	UrlInWriteListKey   = "UrlInWriteListKey"
	UrlInWriteListValue = "YES"
)

type WriteList struct {
	Url    string
	Method string
}

var (
	WriteLists = []WriteList{
		{
			Url:    "/admin/login",
			Method: http.MethodPost,
		},
	}
)

func ParseToken(c *gin.Context) {
	// 请求路径在白名单内
	url := strings.TrimSuffix(strings.TrimPrefix(c.Request.URL.Path, router.ApiV1Prefix), "/")
	for _, w := range WriteLists {
		if w.Url == url && w.Method == c.Request.Method {
			c.Set(UrlInWriteListKey, UrlInWriteListValue)
			c.Next()
			return
		}
	}
	token := c.Request.Header.Get(TokenKeyInHeader)
	// 未找到token
	if token == "" {
		api.Error(c, http.StatusUnauthorized, nil, "身份未认证，请登录")
		return
	}
	// TODO 解析token
	userKey := "123456789"
	c.Set(api.UserIdKey, userKey)
	c.Next()

}

func PermissionCheck(c *gin.Context) {
	// 判断当前请求路径是否为白名单，如果是的话就不需要后面的权限校验，直接跳过
	if value, exists := c.Get(UrlInWriteListKey); exists && value.(string) == UrlInWriteListValue {
		c.Next()
		return
	}
	userKey := api.GetUserKeyFromContext(c)
	// TODO 调用接口确认用户权限
	api.GetRequestLogger(c).Infof("userKey=%v", userKey)
}
