package logger

type Logger interface {
	Info(string)
	Debug(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

type LogServer struct {
	Logger Logger
}

func NewLoggerServer() *LogServer {
	_logger := &LogServer{}
	return _logger
}

func (serv *LogServer) SetLogger(logger Logger) {
	serv.Logger = logger
}
func (serv *LogServer) Info(msg string) {
	serv.Logger.Info(msg)
}
func (serv *LogServer) Debug(msg string) {
	serv.Logger.Debug(msg)
}
func (serv *LogServer) Warn(msg string) {
	serv.Logger.Warn(msg)
}
func (serv *LogServer) Error(msg string) {
	serv.Logger.Error(msg)
}
func (serv *LogServer) Fatal(msg string) {
	serv.Logger.Fatal(msg)
}

// 封装logrus.Fields
// type Fields logrus.Fields

// func SetLevel(level string) {
// 	switch {
// 	case level == "Debug":
// 		logger.SetLevel(logrus.DebugLevel)
// 	case level == "Info":
// 		logger.SetLevel(logrus.InfoLevel)
// 	case level == "Warn":
// 		logger.SetLevel(logrus.WarnLevel)
// 	case level == "Error":
// 		logger.SetLevel(logrus.ErrorLevel)
// 	case level == "Fatal":
// 		logger.SetLevel(logrus.FatalLevel)
// 	}
// }

// func SetLogFormatter(formatter logrus.Formatter) {
// 	logger.Formatter = formatter
// }

// // Debug
// func (serv *LogServer) Debug(info string, msg string) {
// 	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("Failed to log to file, using default stderr")
// 	}
// 	logger.WithFields(logrus.Fields{"info": info}).Debug(msg)
// 	defer file.Close()
// }

// func Info(info string, msg string) {
// 	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("Failed to log to file, using default stderr")
// 	}
// 	logger.WithFields(logrus.Fields{"info": info}).Info(msg)
// 	defer file.Close()
// }

// func Warn(info string, msg string) {
// 	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("Failed to log to file, using default stderr")
// 	}
// 	logger.WithFields(logrus.Fields{"info": info}).Warn(msg)
// 	defer file.Close()
// }

// func Error(info string, msg string) {
// 	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("Failed to log to file, using default stderr")
// 	}
// 	logger.WithFields(logrus.Fields{"info": info}).Error(msg)
// 	defer file.Close()
// }

// func Fatal(info string, msg string) {
// 	file, err := os.OpenFile("../app/logger/files/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("Failed to log to file, using default stderr")
// 	}
// 	logger.WithFields(logrus.Fields{"info": info}).Fatal(msg)
// 	defer file.Close()
// }
