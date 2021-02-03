package log_wrapper

type LOGFUNC func(args ...interface{})
type LOGFUNCT func(template string, args ...interface{})

var Debug LOGFUNC = nil
var Info LOGFUNC = nil
var Warn LOGFUNC = nil
var Error LOGFUNC = nil
var Panic LOGFUNC = nil
var Fatal LOGFUNC = nil

// with template
var Debugf LOGFUNCT = nil
var Infof LOGFUNCT = nil
var Warnf LOGFUNCT = nil
var Errorf LOGFUNCT = nil
var Panicf LOGFUNCT = nil
var Fatalf LOGFUNCT = nil

func init() {
	Debug = zsLog.Debug
	Info = zsLog.Info
	Warn = zsLog.Warn
	Error = zsLog.Error
	Panic = zsLog.Panic
	Fatal = zsLog.Fatal

	Debugf = zsLog.Debugf
	Infof = zsLog.Infof
	Warnf = zsLog.Warnf
	Errorf = zsLog.Errorf
	Panicf = zsLog.Panicf
	Fatalf = zsLog.Fatalf
}

//func Debug(args ...interface{}) {
//	zsLog.Debug(args...)
//}

//func Info(args ...interface{}) {
//	zsLog.Info(args...)
//}

//func Warn(args ...interface{}) {
//	zsLog.Warn(args...)
//}

//func Error(args ...interface{}) {
//	zsLog.Error(args...)
//}

//func Panic(args ...interface{}) {
//	zsLog.Panic(args...)
//}

//func Fatal(args ...interface{}) {
//	zsLog.Fatal(args...)
//}
//
//// with template
//func Debugf(template string, args ...interface{}) {
//	zsLog.Debugf(template, args...)
//}
//
//func Infof(template string, args ...interface{}) {
//	zsLog.Infof(template, args...)
//}
//
//func Warnf(template string, args ...interface{}) {
//	zsLog.Warnf(template, args...)
//}
//
//func Errorf(template string, args ...interface{}) {
//	zsLog.Errorf(template, args...)
//}
//
//func Panicf(template string, args ...interface{}) {
//	zsLog.Panicf(template, args...)
//}
//
//func Fatalf(template string, args ...interface{}) {
//	zsLog.Fatalf(template, args...)
//}
