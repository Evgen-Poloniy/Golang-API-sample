package logger

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// logrus logger implementation
type logrusLogger struct {
	entry *logrus.Entry
	files []*os.File
}

func NewLogrusLogger(config *Config) (Logger, error) {
	l := logrus.New()

	// Set logger format
	switch config.Format {
	case TextFormat:
		l.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		l.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	// Set logger level
	switch config.Level {
	case TraceLevel:
		l.SetLevel(logrus.TraceLevel)
	case DebugLevel:
		l.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		l.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		l.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		l.SetLevel(logrus.ErrorLevel)
	case PanicLevel:
		l.SetLevel(logrus.PanicLevel)
	case FatalLevel:
		l.SetLevel(logrus.FatalLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}

	// Set logger on output stream (Stdout, Stderr, Files)
	writers := make([]io.Writer, 0, len(config.Files))
	files := make([]*os.File, 0, len(config.Files))

	switch config.Output {
	case Stdout:
		writers = append(writers, os.Stdout)
	case Stderr:
		writers = append(writers, os.Stderr)
	default:
		writers = append(writers, os.Stdout)
	}

	for _, filename := range config.Files {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
		}
		writers = append(writers, file)
		files = append(files, file)
	}

	l.SetOutput(io.MultiWriter(writers...))

	return &logrusLogger{
		entry: logrus.NewEntry(l),
		files: files,
	}, nil
}

func (l *logrusLogger) Trace(args ...interface{}) {
	l.entry.Trace(args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l *logrusLogger) Tracef(format string, args ...interface{}) {
	l.entry.Tracef(format, args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.entry.Panicf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogger{entry: l.entry.WithField(key, value)}
}

func (l *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	return &logrusLogger{entry: l.entry.WithFields(fields)}
}

func (l *logrusLogger) WithContext(ctx context.Context) Logger {
	return &logrusLogger{entry: l.entry.WithContext(ctx)}
}

// Closing logger for graceful shutdown
func (l *logrusLogger) Close() error {
	errs := make([]error, 0, len(l.files))

	for _, file := range l.files {
		if err := file.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		return nil
	} else if len(errs) == 1 {
		return fmt.Errorf("error of closing files: %w", errs[0])
	}

	return fmt.Errorf("error of closing files: %w", errors.Join(errs...))
}
