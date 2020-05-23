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
)

var (
	log *zap.Logger // core
)

func init() {
	RegisterWriter()
}

func newLogger() {
	writerCore := zapcore.NewCore(enc, writer, _autoLevel)
	core := zapcore.NewTee(consoleCore, writerCore)
	log = zap.New(core, caller, callerConfig, trace)
}
