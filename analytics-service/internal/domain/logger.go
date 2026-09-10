package domain

type Interface interface {
	Debug(message any, args ...any)
	Info(message any, args ...any)
	Warn(message any, args ...any)
	Error(message any, args ...any)
	Fatal(message any, args ...any)
}
