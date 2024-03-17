package api

import (
	"context"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
)

const (
	RequestIdKey = "request_id"
	LoggerKey    = "_ops-admin-logger-request"
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
	log = logger.ZapLogger.With(RequestIdKey, requestId)
	return log
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(RequestIdKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(RequestIdKey, requestId)
	}
	return requestId
}

func GetUserIdFromContext(c *gin.Context) string {
	if userKey, exists := c.Get(UserIdKey); exists {
		return userKey.(string)
	}
	return ""
}
func GetRequestFromContext(c *gin.Context) string {
	if requestId, exists := c.Get(RequestIdKey); exists {
		return requestId.(string)
	}
	return ""
}
func SetMetadataToContext(c *gin.Context) context.Context {
	userKey := GetUserIdFromContext(c)
	requestId := GetRequestFromContext(c)
	md := metadata.Pairs(UserIdKey, userKey, RequestIdKey, requestId)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}

func GetMetadataFromContext(ctx context.Context) metadata.MD {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	return md
}

func GetRequestIdFromMDContext(ctx context.Context) string {
	md := GetMetadataFromContext(ctx)
	if requestIds, ok := md[strings.ToLower(RequestIdKey)]; ok {
		for _, id := range requestIds {
			return id
		}
	}
	return ""
}

func GetUserIdFromMDContext(ctx context.Context) int {
	md := GetMetadataFromContext(ctx)
	if userIds, ok := md[strings.ToLower(UserIdKey)]; ok {
		for _, id := range userIds {
			userId, _ := strconv.Atoi(id)
			return userId
		}
	}
	return 0
}
