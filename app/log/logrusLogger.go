package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

type Fields logrus.Fields

func NewLogrusLogger() *LogrusLogger {
	var logger = logrus.New()

	_logger := &LogrusLogger{
		logger: logger,
	}
	return _logger
}

func (log *LogrusLogger) SetLevel(level string) {
	// 设置需要些如日志的级别，大于这个级别的日志才会被记录
	switch {
	case level == "Debug":
		log.logger.SetLevel(logrus.DebugLevel)
	case level == "Info":
		log.logger.SetLevel(logrus.InfoLevel)
	case level == "Warn":
		log.logger.SetLevel(logrus.WarnLevel)
	case level == "Error":
		log.logger.SetLevel(logrus.ErrorLevel)
	case level == "Fatal":
		log.logger.SetLevel(logrus.FatalLevel)
	}
}

func (log *LogrusLogger) SetLogFormatter(formatter string) {
	// 设置日志输出的格式，暂定了两个格式，JSON和Text
	switch {
	case formatter == "JSON":
		log.logger.Formatter = &logrus.JSONFormatter{}
	case formatter == "Text":
		log.logger.Formatter = &logrus.TextFormatter{}
	}

}

func (log *LogrusLogger) SetOutput(Type string) {
	// 设置日志的输出方式，暂定两种方式，输出到文件或者是控制台
	switch {
	case Type == "file":
		file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.logger.Out = file
		} else {
			log.logger.Info("Failed to log to file, using default stderr")
		}
	case Type == "os.Stdout":
		log.logger.Out = os.Stdout
	}
}

func (log *LogrusLogger) Info(arg interface{}, msg string) {
	log.logger.WithFields(logrus.Fields{"info": arg}).Info(msg)
}

func (log *LogrusLogger) Debug(arg interface{}, msg string) {
	log.logger.WithFields(logrus.Fields{"info": arg}).Debug(msg)
}
func (log *LogrusLogger) Warn(arg interface{}, msg string) {
	log.logger.WithFields(logrus.Fields{"info": arg}).Warn(msg)
}
func (log *LogrusLogger) Error(arg interface{}, msg string) {
	log.logger.WithFields(logrus.Fields{"info": arg}).Error(msg)

}
func (log *LogrusLogger) Fatal(arg interface{}, msg string) {
	log.logger.WithFields(logrus.Fields{"info": arg}).Fatal(msg)
}
