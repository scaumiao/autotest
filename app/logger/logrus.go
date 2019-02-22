package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// 封装logrus.Fields
type Fields logrus.Fields

func SetLevel(level string) {
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

func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

// Debug
// func Debug(args ...interface{}) {
// 	if logger.Level >= logrus.DebugLevel {
// 		entry := logger.WithFields(logrus.Fields{})
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Debug(args)
// 	}
// }

// Debug
func Debug(info string, msg string) {
	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.WithFields(logrus.Fields{"info": info}).Debug(msg)
	defer file.Close()
}

func Info(info string, msg string) {
	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.WithFields(logrus.Fields{"info": info}).Info(msg)
	defer file.Close()
}

func Warn(info string, msg string) {
	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.WithFields(logrus.Fields{"info": info}).Warn(msg)
	defer file.Close()
}

func Error(info string, msg string) {
	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.WithFields(logrus.Fields{"info": info}).Error(msg)
	defer file.Close()
}

func Fatal(info string, msg string) {
	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.WithFields(logrus.Fields{"info": info}).Fatal(msg)
	defer file.Close()
}

// // 带有field的Debug
// func DebugWithFields(l interface{}, f Fields) {
// 	if logger.Level >= logrus.DebugLevel {
// 		entry := logger.WithFields(logrus.Fields(f))
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Debug(l)
// 	}
// }

// // Info
// // func Info(args ...interface{}) {
// // 	if logger.Level >= logrus.InfoLevel {
// // 		entry := logger.WithFields(logrus.Fields{})
// // 		entry.Data["file"] = fileInfo(2)
// // 		entry.Info(args...)
// // 	}
// // }

// // 带有field的Info
// func InfoWithFields(l interface{}, f Fields) {
// 	if logger.Level >= logrus.InfoLevel {
// 		entry := logger.WithFields(logrus.Fields(f))
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Info(l)
// 	}
// }

// // Warn
// func Warn(args ...interface{}) {
// 	if logger.Level >= logrus.WarnLevel {
// 		entry := logger.WithFields(logrus.Fields{})
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Warn(args...)
// 	}
// }

// // 带有Field的Warn
// func WarnWithFields(l interface{}, f Fields) {
// 	if logger.Level >= logrus.WarnLevel {
// 		entry := logger.WithFields(logrus.Fields(f))
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Warn(l)
// 	}
// }

// // Error
// func Error(args ...interface{}) {
// 	if logger.Level >= logrus.ErrorLevel {
// 		entry := logger.WithFields(logrus.Fields{})
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Error(args...)
// 	}
// }

// // 带有Fields的Error
// func ErrorWithFields(l interface{}, f Fields) {
// 	if logger.Level >= logrus.ErrorLevel {
// 		entry := logger.WithFields(logrus.Fields(f))
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Error(l)
// 	}
// }

// // Fatal
// func Fatal(args ...interface{}) {
// 	if logger.Level >= logrus.FatalLevel {
// 		entry := logger.WithFields(logrus.Fields{})
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Fatal(args...)
// 	}
// }

// // 带有Field的Fatal
// func FatalWithFields(l interface{}, f Fields) {
// 	if logger.Level >= logrus.FatalLevel {
// 		entry := logger.WithFields(logrus.Fields(f))
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Fatal(l)
// 	}
// }

// // Panic
// func Panic(args ...interface{}) {
// 	if logger.Level >= logrus.PanicLevel {
// 		entry := logger.WithFields(logrus.Fields{})
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Panic(args...)
// 	}
// }

// // 带有Field的Panic
// func PanicWithFields(l interface{}, f Fields) {
// 	if logger.Level >= logrus.PanicLevel {
// 		entry := logger.WithFields(logrus.Fields(f))
// 		entry.Data["file"] = fileInfo(2)
// 		entry.Panic(l)
// 	}
// }

// func fileInfo(skip int) string {
// 	_, file, line, ok := runtime.Caller(skip)
// 	if !ok {
// 		file = "../log/files/logrus.log"
// 		line = 1
// 	} else {
// 		slash := strings.LastIndex(file, "/")
// 		if slash >= 0 {
// 			file = file[slash+1:]
// 		}
// 	}
// 	return fmt.Sprintf("%s:%d", file, line)
// }

// func Logger(Fields, msg string) {
// 	file, err := os.OpenFile("../log/files/logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("Failed to log to file, using default stderr")
// 	}
// 	logger.WithFields(logrus.Fields{}).Info(msg)
// 	defer file.Close()
// }
