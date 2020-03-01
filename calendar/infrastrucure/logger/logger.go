package logger

import (
	"github.com/otus/calendar/configs"
	"github.com/otus/calendar/internal/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitZapLogger initialize zap logger.
func InitZapLogger(cfg configs.LoggerConfig) {
	var options []zap.Option

	prodConfig := zap.NewProductionEncoderConfig()
	prodConfig.TimeKey = "T"
	prodConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(prodConfig)
	write := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    50, // megabytes
		MaxBackups: 3,  // old logs
		MaxAge:     3,  // days
		Compress:   true,
	})

	core := zapcore.NewCore(
		encoder,
		write,
		util.GetLogLevel(cfg.LogLevel),
	)

	zap.ReplaceGlobals(zap.New(core, options...))
}
