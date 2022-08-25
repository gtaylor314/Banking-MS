package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// init() is a special Go function that is run only once at startup - here we initialize the zap logger to use throughout the
// application - init() will be called implicitly and is not called explicitly
func init() {
	var err error
	// goal: modify the JSON labels used in zap logger's standard configuration

	// call NewProductionConfig to generate a new zap config - part of config is the EncoderConfig
	config := zap.NewProductionConfig()
	// modify the labels used by the EncoderConfig in config
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Build() constructs a logger - in this case, it writes InfoLevel and above logs to standard error in JSON format (part
	// of the config created by NewProductionConfig()) - we pass in AddCallerSkip(1) which increases the number of callers
	// skipped by caller annotation - since we are wrapping the zap logger only once, we skip one caller so that all logs do
	// not show logger/logger.go as the caller of the zap logger
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// we want to keep the log variable private - we should wrap the zap logger and expose its functionality via our functions

// Info() is a helper function we use to expose zap logger's Info() - the triple dot mmeans that field is a variable argument
// which means we can pass any number of fields to Info()
func Info(message string, fields ...zap.Field) {
	// Info() - zap logger Info() - logs a message at the InfoLevel - message may include fields passed to Info() or
	// fields accumulated on the logger
	log.Info(message, fields...)
}

// Debug() is a helper function we use to expose zap logger's Debug()
func Debug(message string, fields ...zap.Field) {
	// Debug() - zap logger Debug() - logs a message at the DebugLevel
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	// Error() - zap logger Error() - logs a message at the ErrorLevel
	log.Error(message, fields...)
}
