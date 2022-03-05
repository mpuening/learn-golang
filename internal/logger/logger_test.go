package logger

import (
	"testing"
)

var LOG *Logger = NewLogger("test")

func TestLogger(t *testing.T) {

	LOG.Debug("Debug message 1")
	LOG.Debugf("Debug message %d", 2)
	LOG.Info("Info message 1")
	LOG.Infof("Info message %d", 2)
	LOG.Warn("Warn message 1")
	LOG.Warnf("Warn message %d", 2)
	LOG.Fatal("Warn message 1")
	LOG.Fatalf("Warn message %d", 2)
}
