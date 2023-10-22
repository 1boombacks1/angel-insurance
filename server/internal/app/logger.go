package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
		logger.Info("Log to file")
	} else {
		logger.Out = os.Stdout
		logger.Info("Log to stdout")
	}

	return logger
}
