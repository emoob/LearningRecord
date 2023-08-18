package logger

import (
	"testing"
	"time"
)

func TestSetFile(t *testing.T) {
	SetLevel(0)
	SetFile(time.Now().Format(time.DateOnly) + ".log")
	Info("info")
	Debug("debug")
}
