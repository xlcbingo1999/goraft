package utils

import (
	"io"
	"log"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(dir string) *zap.Logger {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	file, err := os.OpenFile(path.Join(dir, "log.txt"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open log file:", err)
	}

	// 设置日志输出的编码格式而已
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.ConsoleSeparator = " | "

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writer := zapcore.AddSync(io.MultiWriter(file, os.Stdout))

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	return zap.New(core, zap.AddCaller())
}
