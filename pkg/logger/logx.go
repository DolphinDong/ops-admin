package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func SetupLogx() {
	logx.MustSetup(logx.LogConf{
		TimeFormat: time.DateTime,
		Mode:       "console",
		Encoding:   "json",
		Level:      "info",
	})
}
