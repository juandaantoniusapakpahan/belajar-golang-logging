package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "sarah").Info("Hello Info")

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{
		"username":   "lusiana",
		"path":       "/dashboard",
		"method":     "GET",
		"statusCode": 200,
	}).Info("Hello Info")
}
