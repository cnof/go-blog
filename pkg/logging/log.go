package logging

import (
	"github.com/Sirupsen/logrus"
	"os"
)

var (
	F *os.File
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(F)
	logrus.SetLevel(logrus.InfoLevel)
}

func MsgInfo(entity *logrus.Entry, msg string) {
	entity.Info(msg)
}
