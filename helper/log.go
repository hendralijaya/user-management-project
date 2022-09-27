package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLog(output string) *logrus.Logger{
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	file, err := os.OpenFile(output, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		PanicIfError(err)
	}
	logger.SetOutput(file)
	return logger
}

