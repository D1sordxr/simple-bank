package interfaces

type Logger interface {
	Info(msg string)
	Infof(msg string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Debug(msg string)
	Debugf(msg string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(msg string)
	Errorf(msg string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panic(msg string)
}
