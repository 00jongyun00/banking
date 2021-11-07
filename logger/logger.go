package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // timestamp 를 좀더 보기 좋게 하기위해서
	encoderConfig.StacktraceKey = ""                      // error 가 발생했을때 stack push 를 비어있게 하기 위해서
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1)) // 설정이 끝나면 build

	//log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
