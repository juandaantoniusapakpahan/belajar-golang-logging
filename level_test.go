package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is a trace level")
	logger.Debug("This is a debug level")
	logger.Info("This is a info level")
	logger.Warn("This is a warn level")
	logger.Error("This is a error level")
}

func TestLoggerLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is a trace level")
	logger.Debug("This is a debug level")
	logger.Info("This is a info level")
	logger.Warn("This is a warn level")
	logger.Error("This is a error level")
}
