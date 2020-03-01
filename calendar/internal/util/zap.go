package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetLogLevel unmarshal text to a zap level notation.
func GetLogLevel(level string) zapcore.Level {

	lvl := zap.DebugLevel
	_ = lvl.UnmarshalText([]byte(level))

	return lvl
}
