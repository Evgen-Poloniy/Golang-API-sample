package logger

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

// slog logger implementation
type slogLogger struct {
	logger *slog.Logger
	files  []*os.File
}

func NewSlogLogger(config *Config) (Logger, error) {
	writers := make([]io.Writer, 0, len(config.Files))
	files := make([]*os.File, 0, len(config.Files))

	// Set logger on output stream (Stdout, Stderr, Files)
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

	writer := io.MultiWriter(writers...)

	// Set logger level
	var level slog.Level
	switch config.Level {
	case TraceLevel:
		level = slog.LevelDebug - 4
	case DebugLevel:
		level = slog.LevelDebug
	case InfoLevel:
		level = slog.LevelInfo
	case WarnLevel:
		level = slog.LevelWarn
	case ErrorLevel:
		level = slog.LevelError
	case PanicLevel:
		level = slog.LevelError + 4
	case FatalLevel:
		level = slog.LevelError + 8
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
	}

	// Set logger format
	switch config.Format {
	case TextFormat:
		handler = slog.NewTextHandler(writer, opts)
	default:
		handler = slog.NewJSONHandler(writer, opts)
	}

	return &slogLogger{
		logger: slog.New(handler),
		files:  files,
	}, nil
}

func (s *slogLogger) Trace(args ...interface{}) {
	s.logger.Log(context.Background(), slog.LevelDebug-4, fmt.Sprint(args...))
}

func (s *slogLogger) Debug(args ...interface{}) {
	s.logger.Debug(fmt.Sprint(args...))
}

func (s *slogLogger) Info(args ...interface{}) {
	s.logger.Info(fmt.Sprint(args...))
}

func (s *slogLogger) Warn(args ...interface{}) {
	s.logger.Warn(fmt.Sprint(args...))
}

func (s *slogLogger) Error(args ...interface{}) {
	s.logger.Error(fmt.Sprint(args...))
}

func (s *slogLogger) Panic(args ...interface{}) {
	msg := fmt.Sprint(args...)
	s.logger.Log(context.Background(), slog.LevelError+4, msg)
	panic(msg)
}

func (s *slogLogger) Fatal(args ...interface{}) {
	s.logger.Log(context.Background(), slog.LevelError+8, fmt.Sprint(args...))
	os.Exit(1)
}

func (s *slogLogger) Tracef(format string, args ...interface{}) {
	s.logger.Log(context.Background(), slog.LevelDebug-4, fmt.Sprintf(format, args...))
}

func (s *slogLogger) Debugf(format string, args ...interface{}) {
	s.logger.Debug(fmt.Sprintf(format, args...))
}

func (s *slogLogger) Infof(format string, args ...interface{}) {
	s.logger.Info(fmt.Sprintf(format, args...))
}

func (s *slogLogger) Warnf(format string, args ...interface{}) {
	s.logger.Warn(fmt.Sprintf(format, args...))
}

func (s *slogLogger) Errorf(format string, args ...interface{}) {
	s.logger.Error(fmt.Sprintf(format, args...))
}

func (s *slogLogger) Panicf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	s.logger.Log(context.Background(), slog.LevelError+4, msg)
	panic(msg)
}

func (s *slogLogger) Fatalf(format string, args ...interface{}) {
	s.logger.Log(context.Background(), slog.LevelError+8, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (s *slogLogger) WithField(key string, value interface{}) Logger {
	return &slogLogger{
		logger: s.logger.With(key, value),
		files:  s.files,
	}
}

func (s *slogLogger) WithFields(fields map[string]interface{}) Logger {
	args := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		args = append(args, k, v)
	}

	return &slogLogger{
		logger: s.logger.With(args...),
		files:  s.files,
	}
}

func (s *slogLogger) WithContext(ctx context.Context) Logger {
	return &slogLogger{
		logger: s.logger.With(slog.Any("context", ctx)),
		files:  s.files,
	}
}

func (s *slogLogger) Close() error {
	errs := make([]error, 0, len(s.files))

	for _, file := range s.files {
		if err := file.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		return nil
	} else if len(errs) == 1 {
		return fmt.Errorf("error closing files: %w", errs[0])
	}

	return fmt.Errorf("error closing files: %w", errors.Join(errs...))
}
