package log

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger(level string) {
	var err error
	var logLevel zap.AtomicLevel

	switch level {
	case "debug":
		logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		logLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		logLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	config := zap.NewProductionConfig()
	config.Level = logLevel
	Logger, err = config.Build()
	if err != nil {
		panic("Không thể khởi tạo logger: " + err.Error())
	}
}

func CloseLogger() {
	if Logger != nil {
		Logger.Sync()
	}
}
