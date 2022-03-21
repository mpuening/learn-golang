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
	LOG.Error("Error message 1")
	LOG.Errorf("Error message %d", 2)
	LOG.Fatal("Fatal message 1")
	LOG.Fatalf("Fatal message %d", 2)
}
