package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type logger struct {
	logger *logrus.Logger
}

type ILogger interface {
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Panic(ctx context.Context, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Trace(ctx context.Context, args ...interface{})
	Tracef(ctx context.Context, format string, args ...interface{})
}

func NewLogger() ILogger {
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	return &logger{logger: logrus.New()}
}

// Debug implements ILogger.
func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Debug(args...)
}

// Debugf implements ILogger.
func (l *logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Debugf(format, args...)
}

// Error implements ILogger.
func (l *logger) Error(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Error(args...)
}

// Errorf implements ILogger.
func (l *logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Errorf(format, args...)
}

// Fatal implements ILogger.
func (l *logger) Fatal(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Fatal(args...)
}

// Fatalf implements ILogger.
func (l *logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Fatalf(format, args...)
}

// Info implements ILogger.
func (l *logger) Info(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Info(args...)
}

// Infof implements ILogger.
func (l *logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Infof(format, args...)
}

// Panic implements ILogger.
func (l *logger) Panic(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Panic(args...)
}

// Panicf implements ILogger.
func (l *logger) Panicf(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Panicf(format, args...)
}

// Trace implements ILogger.
func (l *logger) Trace(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Trace(args...)
}

// Tracef implements ILogger.
func (l *logger) Tracef(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Tracef(format, args...)
}

// Warn implements ILogger.
func (l *logger) Warn(ctx context.Context, args ...interface{}) {
	l.logger.WithContext(ctx).Warn(args...)
}

// Warnf implements ILogger.
func (l *logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Warnf(format, args...)
}
