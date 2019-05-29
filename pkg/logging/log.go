package logging

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(logPath string, logLevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    28,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   false,
	}

	w := zapcore.AddSync(&hook)
	//
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core)

	return logger
}
