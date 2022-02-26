package core

//Logger ...
type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Erro(args ...interface{})
	Errof(format string, args ...interface{})
	Notice(args ...interface{})
	Noticef(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}
