package logging

type Logger interface {
	Debug(msg any, args ...any)
	Info(msg any, args ...any)
	Warn(msg any, args ...any)
	Error(msg any, args ...any)
	Fatal(msg any, args ...any)
}
