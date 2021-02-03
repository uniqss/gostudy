package log_wrapper

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var ZLog *zap.Logger

var zsLog *zap.SugaredLogger

var LogLevelStrs = [...]string{"debug", "info", "warn", "error", "panic", "fatal"}

func GetLogLevel(logLevel string) zapcore.Level {
	for idx, levelStr := range LogLevelStrs {
		if levelStr == logLevel {
			return zapcore.DebugLevel + zapcore.Level(idx)
		}
	}

	fmt.Println("log_wrapper GetLogLevel not recognized. using default zapcore.DebugLevel. logLevel:", logLevel)
	return zapcore.DebugLevel
}

func InitLogger(logDir string, logLevel string, logToConsole bool) {
	//logDir := "../log/"

	var logFilePerm os.FileMode = 0644
	if !strings.HasSuffix(logDir, "/") {
		logDir += "/"
	}
	err := os.MkdirAll(logDir, logFilePerm)
	if err != nil {
		panic("log_wrapper Init_logger os.MkdirAll failed. err:" + err.Error())
	}

	environment := "production"
	//environment := "development"

	srvName := logDir + filepath.Base(os.Args[0])
	pos := strings.Index(srvName, ".exe")
	if pos != -1 {
		srvName = srvName[:pos]
	}
	//	srvName := logdir + srvType + "." + srvID
	var zapLogLevel zapcore.Level = GetLogLevel(logLevel)

	ZLog, _ = NewLogger(environment, srvName, ".log", logToConsole, zapLogLevel)

	zsLog = ZLog.Sugar()
}

func Flush() {
	ZLog.Sync()
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}
func UniqsTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	if 1 == 0 {
		encodeTimeLayout(t, "20060102 15:04:05.000", enc)
	} else {
		encodeTimeLayout(t, "2006-01-02 15:04:05.000", enc)
	}
}

func NewLogger(environment string, logFileName string, suffix string, logToConsole bool, logLevel zapcore.Level) (*zap.Logger, error) {
	if logFileName != "" {
		// create directory if needed.
		pos := strings.LastIndex(logFileName, "/")
		if pos != -1 {
			logFolder := logFileName[0:pos]
			var logFilePerm os.FileMode = 0644

			if logFolder == "" {
				logFolder = "./"
			} else {
				if !strings.HasSuffix(logFolder, "/") {
					logFolder += "/"
				}
				err := os.MkdirAll(logFolder, logFilePerm)
				if err != nil {
					logFileName = "./" + logFileName[pos:]
					fmt.Println("log_wrapper NewLogger os.MkdirAll failed. err:", err, " using current dir instead. logFileName:", logFileName)
				}
			}
		}
	} else {
		// force log to console
		if !logToConsole {
			fmt.Println("log_wrapper NewLogger Warn: zap log NewLogger. as logFileName is empty, logToConsole is set to true")
			logToConsole = true
		}
	}

	priorityConsole := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return logToConsole
	})

	priorityDebug := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel && logLevel <= lev
	})

	priorityInfo := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel && logLevel <= lev
	})

	priorityWarn := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel && logLevel <= lev
	})

	priorityError := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel && logLevel <= lev
	})

	priorityPanic := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.PanicLevel && logLevel <= lev
	})

	priorityFatal := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.FatalLevel && logLevel <= lev
	})

	var prodEncoder zapcore.EncoderConfig
	if environment == "production" {
		prodEncoder = zap.NewProductionEncoderConfig()
	} else {
		prodEncoder = zap.NewDevelopmentEncoderConfig()
	}

	prodEncoder.LevelKey = "LV"
	prodEncoder.TimeKey = "MS"
	prodEncoder.CallerKey = "BT"
	prodEncoder.MessageKey = "MSG"
	prodEncoder.EncodeTime = UniqsTimeEncoder

	var prodEncoderConsole zapcore.EncoderConfig
	prodEncoderConsole = prodEncoder

	if 1 == 1 {
		prodEncoderConsole.EncodeLevel = zapcore.CapitalColorLevelEncoder
		prodEncoder.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		prodEncoderConsole.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		prodEncoder.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	writeSyncerConsole, closeConsole, err := zap.Open("stdout")
	if err != nil {
		closeConsole()
		return nil, err
	}

	MaxSize := 512
	MaxBackups := 1024 * 1024
	MaxAge := 28

	writeSyncerDebug := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName + ".debug" + suffix,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
	})

	writeSyncerInfo := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName + ".info" + suffix,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
	})

	writeSyncerWarn := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName + ".warn" + suffix,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
	})

	writeSyncerError := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName + ".error" + suffix,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
	})

	writeSyncerPanic := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName + ".panic" + suffix,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
	})

	writeSyncerFatal := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName + ".fatal" + suffix,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
	})

	var encoderFunc func(cfg zapcore.EncoderConfig) zapcore.Encoder
	if 1 == 0 {
		// json format
		encoderFunc = zapcore.NewJSONEncoder
	} else {
		// console format
		encoderFunc = zapcore.NewConsoleEncoder
	}

	coreConsole := zapcore.NewCore(encoderFunc(prodEncoderConsole), writeSyncerConsole, priorityConsole)

	coreDebug := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerDebug, priorityDebug)
	coreInfo := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerInfo, priorityInfo)
	coreWarn := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerWarn, priorityWarn)
	coreError := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerError, priorityError)
	corePanic := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerPanic, priorityPanic)
	coreFatal := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerFatal, priorityFatal)

	if logToConsole {
		return zap.New(zapcore.NewTee(coreConsole, coreDebug, coreInfo, coreWarn, coreError, corePanic, coreFatal), zap.AddCaller()), nil
	} else {
		return zap.New(zapcore.NewTee(coreDebug, coreInfo, coreWarn, coreError, corePanic, coreFatal), zap.AddCaller()), nil
	}
}
