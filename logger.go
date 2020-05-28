package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	debug = iota
	info
	warn
	err
	panic
	fatal
)

var (
	log *zap.Logger // core
)

func init() {
	RegisterWriter()
}

func newLogger() {
	if writer != nil {
		writerCore := zapcore.NewCore(enc, writer, _autoLevel)
		core = zapcore.NewTee(core, writerCore)
	}
	if errorWriter != nil {
		writerCore := zapcore.NewCore(enc, errorWriter, zap.ErrorLevel)
		core = zapcore.NewTee(core, writerCore)
	}
	log = zap.New(core, caller, callerConfig, trace).Named(name)
}
