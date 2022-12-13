package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	LogError(err error)
	Fatal(message string, args ...interface{})
}

type Log struct {
	logger *logrus.Entry
}

func New(level string) *Log {
	var l logrus.Level

	switch strings.ToLower(level) {
	case "error":
		l = logrus.ErrorLevel
	case "warm":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}

	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Logger.SetLevel(l)

	return &Log{logger: logger}
}

func (l *Log) Info(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *Log) Debug(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *Log) Warn(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *Log) Error(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *Log) LogError(err error) {
	l.logger.Error(err)
}

func (l *Log) Fatal(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
	os.Exit(1)
}
