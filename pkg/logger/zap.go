package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var (
	ZapLogger *zap.SugaredLogger
)

func SetupZap() {
	ZapLogger = NewLogger("")
}

func NewLogger(logPath string) *zap.SugaredLogger {
	encoder := getEncoder()
	var syncer zapcore.WriteSyncer
	if strings.TrimSpace(logPath) == "" {
		syncer = zapcore.NewMultiWriteSyncer(os.Stdout)
	} else {
		syncer = zapcore.NewMultiWriteSyncer(getWriteSyncer(logPath), os.Stdout)
	}
	core := zapcore.NewCore(encoder, syncer, zap.DebugLevel)
	return zap.New(core, zap.AddCaller()).Sugar()
}
func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "time"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Local().Format("2006-01-02 15:04:05"))
	}
	config.MessageKey = "msg"
	config.CallerKey = "caller"
	config.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(config)
}
func getWriteSyncer(logPath string) zapcore.WriteSyncer {

	logger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    500, // megabytes
		MaxBackups: 7,
		MaxAge:     90,    //days
		Compress:   false, // disabled by default
	}
	return zapcore.AddSync(logger)
}
