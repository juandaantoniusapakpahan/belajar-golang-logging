package belajargolanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Info("Hello Logger")
	fmt.Println("Hello Logger")
}
