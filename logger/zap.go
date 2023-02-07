package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var L *zap.Logger

// Логирование в файл и stdout
func Zap(defaultLogLevel zapcore.Level) *zap.Logger {
	os.MkdirAll("./log", os.ModePerm) //nolint
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFilePath := filepath.Join("./log", time.Now().Format("06.01.02")+".log")
	logFile, _ := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	writer := zapcore.AddSync(logFile)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Info(msg ...any) {
	L.Info(fmt.Sprint(msg)) //nolint
}

func Debug(msg ...any) {
	L.Debug(fmt.Sprint(msg)) //nolint
}

func Error(msg ...any) {
	L.Error(fmt.Sprint(msg)) //nolint
}

func Fatal(msg ...any) {
	L.Fatal(fmt.Sprint(msg)) //nolint
}
