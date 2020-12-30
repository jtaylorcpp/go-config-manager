package types

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
}

func SetLogger(l *logrus.Logger) {
	logger = l
}
