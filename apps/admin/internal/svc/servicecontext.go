package svc

import (
	"context"
	"github.com/DolphinDong/ops-admin/apps/admin/internal/config"
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/DolphinDong/ops-admin/common/models"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		logger: logger.ZapLogger,
		db:     models.GetDB(),
	}
}

func (svc ServiceContext) Logger(ctx context.Context) *zap.SugaredLogger {
	requestId := api.GetRequestIdFromMDContext(ctx)
	if requestId != "" {
		return svc.logger.With(api.RequestIdKey, requestId)
	}
	svc.logger.Warnf("Request Id not found in context")
	return svc.logger
}

func (svc ServiceContext) GetUserId(ctx context.Context) int {
	return api.GetUserIdFromMDContext(ctx)
}

func (svc ServiceContext) GetDB(ctx context.Context) *gorm.DB {
	return svc.db.WithContext(ctx)
}
