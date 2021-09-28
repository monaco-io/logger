package logger

import (
	"google.golang.org/grpc/metadata"
)

func _handler(c *contextLogger, level int, msg string, keyValues ...interface{}) {
	sugar := log.Sugar()
	if c != nil {
		var in bool
		imd, iok := metadata.FromIncomingContext(c.Context)
		if iok && len(imd[xRequestID]) > 0 {
			sugar = sugar.With(xRequestID, imd[xRequestID][0])
			in = true
		}
		if !in {
			omd, ook := metadata.FromOutgoingContext(c.Context)
			if ook && len(omd[xRequestID]) > 0 {
				sugar = sugar.With(xRequestID, omd[xRequestID][0])
			}
		}
	}
	defer func() { _ = log.Sync() }()

	switch level {
	case debug:
		sugar.Debugw(msg, keyValues...)
	case info:
		sugar.Infow(msg, keyValues...)
	case warn:
		sugar.Warnw(msg, keyValues...)
	case err:
		sugar.Errorw(msg, keyValues...)
	case panic:
		sugar.Panicw(msg, keyValues...)
	case fatal:
		sugar.Fatalw(msg, keyValues...)
	case dev:
		sugar.Infow(msg, keyValues...)
	}
}
