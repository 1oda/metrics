package log

import (
	"log"
	"metrics/utils/config"
	"os"
	"path"
	"strconv"

	logs "github.com/sirupsen/logrus"
)

const (
	InfoLevel = 4
	WarnLevel = 3
)

var logger *logs.Logger

func InitLog() {
	logName := path.Join(config.AppConfBO.LogPath, config.AppConfBO.LogFile)
	_, err := os.Stat(logName)
	if err != nil {
		err = os.WriteFile(logName, []byte(""), os.ModePerm)
	}

	f, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	logLevel, err := strconv.Atoi(config.AppConfBO.LogLevel)
	if err != nil {

	}
	if logLevel == 0 {
		logLevel = InfoLevel
	}

	logger = logs.New()
	logger.SetFormatter(&logs.JSONFormatter{})
	logger.SetOutput(f)
	logger.SetLevel(logs.Level(logLevel))
}

func Info(v interface{}) {
	logger.Info(v)
}

func Error(v interface{}) {
	logger.Error(v)
}

func Warn(v interface{}) {
	logger.Warn(v)
}
