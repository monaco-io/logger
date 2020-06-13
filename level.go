package logger

func D(msg string, keyValues ...interface{}) {
	_handler(debug, msg, keyValues...)
}

func I(msg string, keyValues ...interface{}) {
	_handler(info, msg, keyValues...)
}

func W(msg string, keyValues ...interface{}) {
	_handler(warn, msg, keyValues...)
}

func E(msg string, keyValues ...interface{}) {
	_handler(err, msg, keyValues...)
}

func P(msg string, keyValues ...interface{}) {
	_handler(panic, msg, keyValues...)
}

func F(msg string, keyValues ...interface{}) {
	_handler(fatal, msg, keyValues...)
}

func T(keyValues ...interface{}) {
	_handler(dev, "dev", keyValues...)
}
