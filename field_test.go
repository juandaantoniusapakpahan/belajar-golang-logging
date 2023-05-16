package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "jhonson").Info("Hello Info")
	logger.WithField("username", "lisa").WithField("name", "Lisa").Info("Hello Info")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "Jhonson",
		"mehtod":   "GET",
		"path":     "/books",
		"status":   200,
	}).Info("Hello Info")
}
