package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// 封装logrus.Fields
type Fields logrus.Fields

func SetLevel(level string) {
	// 设置需要些如日志的级别，大于这个级别的日志才会被记录
	switch {
	case level == "Debug":
		logger.SetLevel(logrus.DebugLevel)
	case level == "Info":
		logger.SetLevel(logrus.InfoLevel)
	case level == "Warn":
		logger.SetLevel(logrus.WarnLevel)
	case level == "Error":
		logger.SetLevel(logrus.ErrorLevel)
	case level == "Fatal":
		logger.SetLevel(logrus.FatalLevel)
	}
}

func SetLogFormatter(formatter string) {
	// 设置日志输出的格式，暂定了两个格式，JSON和Text
	switch {
	case formatter == "JSON":
		logger.Formatter = &logrus.JSONFormatter{}
	case formatter == "Text":
		logger.Formatter = &logrus.TextFormatter{}
	}

}

func SetOutput(Type string) {
	// 设置日志的输出方式，暂定两种方式，输出到文件或者是控制台
	switch {
	case Type == "file":
		file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.Out = file
		} else {
			logger.Info("Failed to log to file, using default stderr")
		}
	case Type == "os.Stdout":
		logger.Out = os.Stdout
	}
}

func init() {
	// 初始化
	SetLevel("Debug")
	SetLogFormatter("JSON")
	SetOutput("file")
}

func Debug(info string, msg string) {
	logger.WithFields(logrus.Fields{"info": info}).Debug(msg)
}

func Info(info string, msg string) {
	logger.WithFields(logrus.Fields{"info": info}).Info(msg)
}

func Warn(info string, msg string) {
	logger.WithFields(logrus.Fields{"info": info}).Warn(msg)
}

func Error(info string, msg string) {
	logger.WithFields(logrus.Fields{"info": info}).Error(msg)
}

func Fatal(info string, msg string) {
	logger.WithFields(logrus.Fields{"info": info}).Fatal(msg)
}
