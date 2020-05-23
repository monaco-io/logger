package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var (
	_encoderConfig = zap.NewDevelopmentEncoderConfig()
	_autoLevel     = zap.NewAtomicLevel()

	writer       zapcore.WriteSyncer
	caller       = zap.AddCaller()
	callerConfig = zap.AddCallerSkip(2)
	enc          = zapcore.NewJSONEncoder(_encoderConfig)
	trace        = zap.AddStacktrace(zap.ErrorLevel)

	consoleCore = zapcore.NewCore(zapcore.NewConsoleEncoder(_encoderConfig), os.Stdout, _autoLevel)
)

func RegisterWriter(writers ...io.Writer) {
	var writerTmp zapcore.WriteSyncer
	for _, w := range writers {
		if writerTmp == nil {
			writerTmp = zapcore.AddSync(w)
			break
		}
		writerTmp = zap.CombineWriteSyncers(writerTmp, zapcore.AddSync(w))
	}
	writer = writerTmp
	newLogger()
}
