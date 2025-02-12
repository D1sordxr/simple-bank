package logger

import (
	"log/slog"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(logger *slog.Logger) *Logger {
	return &Logger{Logger: logger}
}

func (l *Logger) Info(msg string, attrs ...any) {
	l.Logger.Info(msg, attrs...)
}

func (l *Logger) Error(msg string, attrs ...any) {
	l.Logger.Error(msg, attrs...)
}

func (l *Logger) Debug(msg string, attrs ...any) {
	l.Logger.Debug(msg, attrs...)
}

func (l *Logger) String(key string, value string) slog.Attr {
	return slog.String(key, value)
}

func (l *Logger) Float64(key string, v float64) slog.Attr {
	return slog.Float64(key, v)
}

func (l *Logger) Group(key string, args ...any) slog.Attr {
	return slog.Group(key, args...)
}
