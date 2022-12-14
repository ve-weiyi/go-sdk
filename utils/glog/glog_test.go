package glog

import (
	"go.uber.org/zap"
	"testing"
)

func TestName(t *testing.T) {
	Error("hello")
	Debug("123")
	Info("info")
	ReplaceZapGlobals()
	zap.L().Debug("hello world")
}
