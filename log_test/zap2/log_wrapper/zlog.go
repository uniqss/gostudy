package log_wrapper

import (
	"go.uber.org/zap"
)

var zsLog *zap.SugaredLogger

func init() {
	zLog := ZLog
	zsLog = zLog.Sugar()
}

func Debug(args ...interface{}) {
	zsLog.Debug(args...)
}

func Info(args ...interface{}) {
	zsLog.Info(args...)
}

func Warn(args ...interface{}) {
	zsLog.Warn(args...)
}

func Error(args ...interface{}) {
	zsLog.Error(args...)
}

func Panic(args ...interface{}) {
	zsLog.Panic(args...)
}

func Fatal(args ...interface{}) {
	zsLog.Fatal(args...)
}

// with template
func Debugf(template string, args ...interface{}) {
	zsLog.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	zsLog.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	zsLog.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	zsLog.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	zsLog.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	zsLog.Fatalf(template, args...)
}
