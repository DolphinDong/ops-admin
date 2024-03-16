package middleware

import (
	"fmt"
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 打印请求信息
func PrintRequestInfo(c *gin.Context) {
	ip := c.ClientIP()
	url := c.Request.URL
	startTime := time.Now().UnixMicro()
	c.Next()
	endTime := time.Now().UnixMicro()
	logger.ZapLogger.Infow("", api.TrafficKey, api.GenerateMsgIDFromContext(c),
		"method", c.Request.Method,
		"uri", url.RequestURI(),
		"client_ip", ip,
		"code", c.Writer.Status(),
		"response_time", fmt.Sprintf("%vms", float32(endTime-startTime)/1000))
}

func NoRouter(c *gin.Context) {
	ip := c.ClientIP()
	url := c.Request.URL
	api.GetRequestLogger(c).Infow("Resource not found",
		api.TrafficKey, api.GenerateMsgIDFromContext(c),
		"method", c.Request.Method,
		"uri", url.RequestURI(),
		"client_ip", ip,
		"code", http.StatusNotFound)
	api.Error(c, http.StatusNotFound, nil, "Resource not found")
}
