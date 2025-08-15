package untils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

func SetUpLogger() (*zap.Logger, error) {
	logger, _ := zap.NewProduction()
	//创建日志目录(如果不存在)
	logDir := "./log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, 0755); err != nil {
			return nil, fmt.Errorf("创建日志目录失败:%v", err)
		}
	}
	log.Println("日志目录创建成功")
	//生成带日期的文件
	dTime := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("../logs/%s.log", dTime)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    100,
		MaxAge:     3,
		MaxBackups: 28,
		Compress:   true,
	})

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		zapcore.DebugLevel,
	)

	logger = zap.New(core)
	return logger, nil
}
