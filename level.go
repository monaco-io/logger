package logger

func I(msg string, keyValues ...interface{}) {
	_handler(info, msg, keyValues...)
}

func D(msg string, keyValues ...interface{}) {
	_handler(debug, msg, keyValues...)
}

func W(msg string, keyValues ...interface{}) {
	_handler(warn, msg, keyValues...)
}

func E(msg string, keyValues ...interface{}) {
	_handler(err, msg, keyValues...)
}
