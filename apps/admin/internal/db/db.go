package db

import (
	"github.com/DolphinDong/ops-admin/common/models/admin"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		new(admin.User),
	)
	if err != nil {
		logger.ZapLogger.Fatalf("Migrate db failed: %+v", errors.WithStack(err))
	}
	logger.ZapLogger.Info("Migrate db success")
}
