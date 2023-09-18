package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"

	"poker/pkg/config"
)

func InitLogger() error {
	setLevel()

	setFormatter()

	if err := setOutput(); err != nil {
		return err
	}

	return nil
}

func setLevel() {
	level, err := logrus.ParseLevel(config.GetLoggerLevel())
	if err != nil {
		level = logrus.InfoLevel
	}

	logrus.SetLevel(level)
}

func setOutput() error {
	output := config.GetLoggerOutput()

	if output != "stdout" {
		f, err := os.OpenFile(path.Join(output, fmt.Sprintf("poker-%s.log", time.Now().Format("2006-01-02T15:04:05"))), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

func setFormatter() {
	if config.GetLoggerFormat() == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
