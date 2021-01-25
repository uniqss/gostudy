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

func init() {
	logDir := "./log/"
	logLevel := "debug"

	environment := "production"
	//environment := "development"

	_ = logLevel

	srvName := logDir + filepath.Base(os.Args[0])
	pos := strings.Index(srvName, ".exe")
	if pos != -1 {
		srvName = srvName[:pos]
	}
	//	srvName := logdir + srvType + "." + srvID
	ZLog, _ = NewLogger(environment, srvName, ".log", true)
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
	encodeTimeLayout(t, "20060102 15:04:05.000", enc)
}

func NewLogger(environment string, logFileName string, suffix string, logToConsole bool) (*zap.Logger, error) {
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
					fmt.Println("log_wrapper InitLog os.MkdirAll failed. err:", err, " using current dir instead. logFileName:", logFileName)
				}
			}
		}
	} else {
		// force log to console
		if !logToConsole {
			fmt.Println("Warn: zap log NewLogger. as logFileName is empty, logToConsole is set to true")
		}
	}

	priorityConsole := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return logToConsole
	})

	priorityDebug := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})

	priorityInfo := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})

	priorityWarn := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})

	priorityError := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel
	})

	priorityPanic := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.PanicLevel
	})

	priorityFatal := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.FatalLevel
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

	coreConsole := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerConsole, priorityConsole)

	coreDebug := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerDebug, priorityDebug)
	coreInfo := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerInfo, priorityInfo)
	coreWarn := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerWarn, priorityWarn)
	coreError := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerError, priorityError)
	corePanic := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerPanic, priorityPanic)
	coreFatal := zapcore.NewCore(encoderFunc(prodEncoder), writeSyncerFatal, priorityFatal)

	return zap.New(zapcore.NewTee(coreConsole, coreDebug, coreInfo, coreWarn, coreError, corePanic, coreFatal), zap.AddCaller()), nil
}
