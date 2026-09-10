package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(level string) *Logger {
	var l zerolog.Level
	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "fatal":
		l = zerolog.FatalLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)
	skipFrameCount := 4
	logger := zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
		Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) log(level zerolog.Level, message string, args ...any) {
	if len(args) == 0 {
		l.logger.WithLevel(level).Msg(message)
	} else {
		l.logger.WithLevel(level).Msgf(message, args...)
	}
}

func (l *Logger) msg(level zerolog.Level, message any, args ...any) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(level, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

func (l *Logger) Info(message any, args ...any) {
	l.msg(zerolog.InfoLevel, message, args...)
}

func (l *Logger) Error(message any, args ...any) {
	l.msg(zerolog.ErrorLevel, message, args...)
}

func (l *Logger) Fatal(message any, args ...any) {
	l.msg(zerolog.FatalLevel, message, args...)
}

func (l *Logger) Debug(message any, args ...any) {
	l.msg(zerolog.DebugLevel, message, args...)
}

func (l *Logger) Warn(message any, args ...any) {
	l.msg(zerolog.WarnLevel, message, args...)
}
