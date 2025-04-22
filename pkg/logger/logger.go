// Package logger defines a custom logger interface and implements it using zap.
package logger

import (
	"context"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

// Logger is a custom logger interface
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

// ZapLogger implements the Logger interface using zap
type ZapLogger struct {
	logger *zap.SugaredLogger
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level  string
	Format string
	Output string
}

// NewLogger creates a new zap logger instance
func NewLogger(cfg LogConfig) *ZapLogger {
	// Configure encoder
	var encoder zapcore.Encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	switch cfg.Format {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// Configure output
	var output io.Writer
	switch cfg.Output {
	case "stderr":
		output = os.Stderr
	default:
		output = os.Stdout
	}

	// Create core
	atomicLevel := getLogLevel(cfg.Level)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(output),
		atomicLevel,
	)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &ZapLogger{
		logger: logger.Sugar(),
	}
}

// NewLoggerFromConfig creates a logger from the standard config
func NewLoggerFromConfig(level, format, output string) *ZapLogger {
	return NewLogger(LogConfig{
		Level:  level,
		Format: format,
		Output: output,
	})
}

// getLogLevel converts string level to zapcore.Level
func getLogLevel(level string) zap.AtomicLevel {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevelAt(zapLevel)
	return atomicLevel
}

// Debug logs debug messages
func (l *ZapLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}

// Info logs info messages
func (l *ZapLogger) Info(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

// Warn logs warning messages
func (l *ZapLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

// Error logs error messages
func (l *ZapLogger) Error(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

// Fatal logs fatal messages and exits
func (l *ZapLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}

// GormLoggerAdapter adapts our logger to work with GORM
type GormLoggerAdapter struct {
	Logger   Logger
	LogLevel gormlogger.LogLevel
}

// LogMode implementation for GORM
func (l *GormLoggerAdapter) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs info messages for GORM
func (l *GormLoggerAdapter) Info(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Info(msg, args...)
}

// Warn logs warning messages for GORM
func (l *GormLoggerAdapter) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Warn(msg, args...)
}

// Error logs error messages for GORM
func (l *GormLoggerAdapter) Error(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Error(msg, args...)
}

// Trace logs SQL queries for GORM
func (l *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		l.Logger.Error("SQL Error: %s, Elapsed: %v, Rows: %d, Error: %v", sql, elapsed, rows, err)
		return
	}

	if elapsed > 200*time.Millisecond {
		l.Logger.Warn("Slow SQL: %s, Elapsed: %v, Rows: %d", sql, elapsed, rows)
		return
	}

	if l.LogLevel <= gormlogger.Info { // Debug level
		l.Logger.Debug("SQL: %s, Elapsed: %v, Rows: %d", sql, elapsed, rows)
	}
}
