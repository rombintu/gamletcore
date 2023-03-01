package tools

import (
	"os"

	logrus "github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	return logrus.New()
}
