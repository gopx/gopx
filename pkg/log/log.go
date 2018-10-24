package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{}
	if env, ok := os.LookupEnv("DEBUG"); ok && env == "1" {
		log.Level = logrus.DebugLevel
	} else {
		log.Level = logrus.InfoLevel
	}
}

func callStackInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

// Debug logs a message with call stack trace at level Debug on the standard logger.
func Debug(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"TRACE": callStackInfo(),
	}).Debugf(format, args...)
}

// Info logs a message with call stack trace at level Info on the standard logger.
func Info(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"TRACE": callStackInfo(),
	}).Infof(format, args...)
}

// Warn logs a message with call stack trace at level Warn on the standard logger.
func Warn(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"TRACE": callStackInfo(),
	}).Warnf(format, args...)
}

// Error logs a message with call stack trace at level Error on the standard logger.
func Error(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"TRACE": callStackInfo(),
	}).Errorf(format, args...)
}

// Fatal logs a message with call stack trace at level Fatal on the standard logger.
func Fatal(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"TRACE": callStackInfo(),
	}).Fatalf(format, args...)
}

// Panic logs a message with call stack trace at level Panic on the standard logger.
func Panic(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"TRACE": callStackInfo(),
	}).Panicf(format, args...)
}
