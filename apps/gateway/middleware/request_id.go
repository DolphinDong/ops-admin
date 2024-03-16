package middleware

import (
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

// RequestId 自动增加requestId
func RequestId(c *gin.Context) {
	requestId := c.GetHeader(api.TrafficKey)
	if requestId == "" {
		requestId = c.GetHeader(strings.ToLower(api.TrafficKey))
	}
	if requestId == "" {
		requestId = uuid.New().String()
	}
	c.Request.Header.Set(api.TrafficKey, requestId)
	c.Request.Header.Set(api.TrafficKey, requestId)
	c.Set(api.TrafficKey, requestId)
	c.Set(api.LoggerKey, logger.ZapLogger.With(api.TrafficKey, requestId))
	c.Next()
}
