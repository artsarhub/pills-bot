package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

// init Инициализация логгера для использования его во всем приложении.
// init будет выполнен один раз, независимо от количества импортов в разных местах приложения.
func init() {
	// Инициализация для режима разработки.
	localLogger, err := zap.NewDevelopment()
	// Вариант инициализации для продакшена.
	//localLogger, err := zap.NewProduction()

	if err != nil {
		log.Fatal("Ошибка инициализации логгера zap", err)
	}

	logger = localLogger

}

func Fatal(msg string, keysAndValues ...interface{}) {
	sugar := logger.Sugar()
	sugar.Fatalw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	sugar := logger.Sugar()
	sugar.Errorw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	sugar := logger.Sugar()
	sugar.Warnw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	sugar := logger.Sugar()
	sugar.Infow(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	sugar := logger.Sugar()
	sugar.Debugw(msg, keysAndValues...)
}

func DebugZap(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
