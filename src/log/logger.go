package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Init(path string) {
	// logFile, _ := os.Create(path)

	lumberjack := &lumberjack.Logger{
		Filename:   path + "/log.txt",
		MaxSize:    10000000,
		MaxAge:     10,
		MaxBackups: 3,
	}
	ws := zapcore.AddSync(lumberjack)
	enCfg := zap.NewProductionEncoderConfig()
	enCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(enCfg)
	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)
	logger = zap.New(core)
}

func Debug(info string) {
	logger.Debug(info)
}

func Debugf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	logger.Debug(msg)
}

func Error(info string) {
	logger.Error(info)
}

func Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	logger.Error(msg)
}
