package logger

func _handler(level int, msg string, keyValues ...interface{}) {
	defer func() { _ = log.Sync() }()
	switch level {
	case debug:
		log.Sugar().Debugw(msg, keyValues...)
	case info:
		log.Sugar().Infow(msg, keyValues...)
	case warn:
		log.Sugar().Warnw(msg, keyValues...)
	case err:
		log.Sugar().Errorw(msg, keyValues...)
	}
}
