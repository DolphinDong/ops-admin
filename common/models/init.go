package models

import (
	"fmt"
	"github.com/DolphinDong/ops-admin/common/config"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	_db   *gorm.DB
	DSN   = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	_once = sync.Once{}
)

type dbLogger struct {
}

func (l *dbLogger) Printf(msg string, args ...interface{}) {
	logger.ZapLogger.Infof(msg, args...)
}

func SetupDB(mysqlConf config.Mysql) {
	_once.Do(func() {
		newLogger := glog.New(
			&dbLogger{}, // io writer
			glog.Config{
				SlowThreshold: time.Second,                           // Slow SQL threshold
				LogLevel:      getLoggerLevel(mysqlConf.SqlLogLevel), // Log level
			},
		)
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf(DSN, mysqlConf.Username, mysqlConf.Password, mysqlConf.HostName, mysqlConf.Port, mysqlConf.Database), // DSN data source name 			// 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			logger.ZapLogger.Fatalf("Init database connection failed: %+v", errors.WithStack(err))
		}
		sqlDB, err := db.DB()
		if err != nil {
			logger.ZapLogger.Fatalf("Setup database pool failed: %+v", errors.WithStack(err))
		}
		sqlDB.SetMaxIdleConns(mysqlConf.MaxIdle)
		sqlDB.SetMaxOpenConns(mysqlConf.MaxOpen)
		sqlDB.SetConnMaxLifetime(time.Hour)

		_db = db
	})
	logger.ZapLogger.Info("Init database connection success")
}

func GetDB() *gorm.DB {
	return _db
}

func getLoggerLevel(level string) glog.LogLevel {
	switch level {
	case "info":
		return glog.Info
	case "warn":
		return glog.Warn
	case "error":
		return glog.Error
	case "silent":
		return glog.Silent
	default:
		return glog.Warn
	}
	return glog.Warn
}
