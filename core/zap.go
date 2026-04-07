package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 初始化日志
func Zap() *zap.Logger {
	// 日志配置
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式

	// 构建日志实例
	logger, err := config.Build()
	if err != nil {
		panic("初始化 zap 日志失败: " + err.Error())
	}

	// 替换全局日志
	defer logger.Sync()
	return logger
}
