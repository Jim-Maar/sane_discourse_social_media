package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	// Set output to stdout
	Log.SetOutput(os.Stdout)

	// Set log level (can be configured via environment variable)
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	// Use JSON formatter for structured logging
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})

	Log.Info("Logger initialized")
}

// GetLogger returns the configured logger instance
func GetLogger() *logrus.Logger {
	if Log == nil {
		Init()
	}
	return Log
}
