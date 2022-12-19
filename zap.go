package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	fileName := "micro.log"
	zapSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName, // 文件名
		MaxSize:    512,      // 文件大小
		MaxBackups: 0,        // 最大备份数
		LocalTime:  true,
		Compress:   true, // 是否启用压缩
	})
	// 编码
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 核心配置
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapSync, zap.NewAtomicLevelAt(zap.DebugLevel))
	// 创建logger
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// 创建语法糖
	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(temp string, args ...interface{}) {
	logger.Debugf(temp, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(temp string, args ...interface{}) {
	logger.Infof(temp, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(temp string, args ...interface{}) {
	logger.Warnf(temp, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(temp string, args ...interface{}) {
	logger.Errorf(temp, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(temp string, args ...interface{}) {
	logger.Fatalf(temp, args...)
}
