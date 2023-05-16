package belajargolanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampelHook struct {
}

func (hook *SampelHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (hook *SampelHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Simple Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampelHook{})
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Hello Info")
	logger.Debug("Hello debug")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")
}
