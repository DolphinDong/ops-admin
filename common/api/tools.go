package api

import (
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	TrafficKey = "request_id"
	LoggerKey  = "_ops-admin-logger-request"
)

const (
	UserIdKey = "USER_KEY"
)

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *zap.SugaredLogger {
	var log *zap.SugaredLogger
	l, ok := c.Get(LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*zap.SugaredLogger)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := GenerateMsgIDFromContext(c)
	log = logger.ZapLogger.With(TrafficKey, requestId)
	return log
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

func GetUserKeyFromContext(c *gin.Context) string {
	if userKey, exists := c.Get(UserIdKey); exists {
		return userKey.(string)
	}
	return ""
}
