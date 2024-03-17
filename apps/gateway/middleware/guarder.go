package middleware

import (
	"github.com/DolphinDong/ops-admin/apps/gateway/router"
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/DolphinDong/ops-admin/common/rpc"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	TokenKeyInHeader    = "Authorization"
	UrlInWriteListKey   = "UrlInWriteListKey"
	UrlInWriteListValue = "YES"
)

const ()

type WriteList struct {
	Url    string
	Method string
}

var (
	WriteLists = []WriteList{
		{
			Url:    router.ApiV1Prefix + "/admin/login",
			Method: http.MethodPost,
		},
	}
)

func ParseToken(c *gin.Context) {
	// 请求路径在白名单内
	url := strings.TrimSuffix(c.Request.URL.Path, "/")
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
	// 解析token
	checkTokenRes, err := rpc.AdminClient.CheckToken(api.SetMetadataToContext(c), &pb.CheckTokenReq{Token: strings.TrimPrefix(token, "Bearer ")})
	if err != nil {
		logger.ZapLogger.Warnf("Token解析失败: %v", err)
		api.Error(c, http.StatusUnauthorized, nil, "Token解析失败")
		return
	}
	// 解析不成功
	if !checkTokenRes.Success {
		logger.ZapLogger.Warnf(checkTokenRes.Message)
		api.Error(c, http.StatusUnauthorized, nil, checkTokenRes.Message)
		return
	}

	c.Set(api.UserIdKey, strconv.Itoa(int(checkTokenRes.UserId)))
	c.Next()
}

func PermissionCheck(c *gin.Context) {
	// 判断当前请求路径是否为白名单，如果是的话就不需要后面的权限校验，直接跳过
	if value, exists := c.Get(UrlInWriteListKey); exists && value.(string) == UrlInWriteListValue {
		c.Next()
		return
	}
	userIdStr := api.GetUserIdFromContext(c)
	userId, _ := strconv.Atoi(userIdStr)
	checkPermissionRes, err := rpc.AdminClient.CheckPermission(api.SetMetadataToContext(c), &pb.CheckPermissionReq{
		Url:    c.Request.URL.Path,
		Method: c.Request.Method,
		UserId: int64(userId),
	})
	if err != nil {
		logger.ZapLogger.Warnf("权限校验异常: %v", err)
		api.Error(c, http.StatusForbidden, nil, "权限校验异常")
		return
	}
	if !checkPermissionRes.Success {
		logger.ZapLogger.Warnf("权限不足")
		api.Error(c, http.StatusForbidden, nil, "权限不足")
		return
	}
	c.Next()
}
