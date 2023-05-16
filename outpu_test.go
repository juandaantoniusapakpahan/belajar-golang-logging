package belajargolanglogging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutpu(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}
