package logger

import (
	"aws-iot-jobs-tui/where"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func Setup() error {
	
	// if logs dir does not exist, where.Logs() creates it
	logPath := where.Logs()
	
	logFile, err := os.OpenFile(filepath.Join(logPath, "iotjob.log"), os.O_CREATE |os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open iotjob.log file, here's why: %v", err)
	}

	// Set TraceLevel as base logging level.
	// Log everything called by:
	/*
		logger.Trace()
		logger.Debug()
		logger.Info()
		logger.Warn()
		logger.Error()
		logger.Fatal()
		logger.Panic()
	*/
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logrus.SetLevel(logrus.TraceLevel)

	return nil
}

// Wrap logrus methods for logging
func Panic(args ...any) {
	logrus.Panic(args...)
}

func Panicf(format string, args ...any) {
	logrus.Panicf(format, args...)
}

func Fatal(args ...any) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...any) {
	logrus.Fatalf(format, args...)
}

func Error(args ...any) {
	logrus.Error(args...)
}

func Errorf(format string, args ...any) {
	logrus.Errorf(format, args...)
}

func Warn(args ...any) {
	logrus.Warn(args...)
}

func Warnf(format string, args ...any) {
	logrus.Warnf(format, args...)
}

func Info(args ...any) {
	logrus.Info(args...)
}

func Infof(format string, args ...any) {
	logrus.Infof(format, args...)
}

func Debug(args ...any) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...any) {
	logrus.Debugf(format, args...)
}

func Trace(args ...any) {
	logrus.Trace(args...)
}

func Tracef(format string, args ...any) {
	logrus.Tracef(format, args...)
}
