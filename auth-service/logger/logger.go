package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger
var LoggingEntry *logrus.Entry

func InitLogger()  {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetReportCaller(true)

	Logger.SetOutput(&lumberjack.Logger{
		Filename:   "./logger/logs/auth-service.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     30, //days
	})
}

