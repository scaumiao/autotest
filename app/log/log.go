package logger

type Logger interface {
	Info(interface{}, string)
	Debug(interface{}, string)
	Warn(interface{}, string)
	Error(interface{}, string)
	Fatal(interface{}, string)
	SetLevel(string)
	SetLogFormatter(string)
	SetOutput(string)
	NewLogrusLogger()
}

type LogServer struct {
	Logger Logger
}

func (serv *LogServer) init() {
	serv.Logger.SetLevel("Info")
	serv.Logger.SetLogFormatter("JSON")
	serv.Logger.SetOutput("file")
}

func NewLoggerServer() *LogServer {
	_logger := &LogServer{}
	return _logger
}

func (serv *LogServer) SetLogger(logger Logger) {
	serv.Logger = logger
}
func (serv *LogServer) Info(arg interface{}, msg string) {
	serv.Logger.Info(arg, msg)
}
func (serv *LogServer) Debug(arg interface{}, msg string) {
	serv.Logger.Debug(arg, msg)
}
func (serv *LogServer) Warn(arg interface{}, msg string) {
	serv.Logger.Warn(arg, msg)
}
func (serv *LogServer) Error(arg interface{}, msg string) {
	serv.Logger.Error(arg, msg)
}
func (serv *LogServer) Fatal(arg interface{}, msg string) {
	serv.Logger.Fatal(arg, msg)
}

func (serv *LogServer) SetLevel(level string) {
	serv.Logger.SetLevel(level)
}

func (serv *LogServer) SetLogFormatter(formatter string) {
	serv.Logger.SetLogFormatter(formatter)
}

func (serv *LogServer) SetOutput(Type string) {
	serv.Logger.SetOutput(Type)
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
