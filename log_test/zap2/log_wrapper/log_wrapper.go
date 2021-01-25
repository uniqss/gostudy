package log_wrapper

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	writeSyncerDebug, closeDebug, err := zap.Open(logFileName + ".debug" + suffix)
	if err != nil {
		closeDebug()
		return nil, err
	}

	writeSyncerInfo, closeInfo, err := zap.Open(logFileName + ".info" + suffix)
	if err != nil {
		closeInfo()
		return nil, err
	}

	writeSyncerWarn, closeWarn, err := zap.Open(logFileName + ".warn" + suffix)
	if err != nil {
		closeWarn()
		return nil, err
	}

	writeSyncerError, closeError, err := zap.Open(logFileName + ".error" + suffix)
	if err != nil {
		closeError()
		return nil, err
	}

	writeSyncerPanic, closePanic, err := zap.Open(logFileName + ".panic" + suffix)
	if err != nil {
		closePanic()
		return nil, err
	}

	writeSyncerFatal, closeFatal, err := zap.Open(logFileName + ".fatal" + suffix)
	if err != nil {
		closeFatal()
		return nil, err
	}

	coreConsole := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerConsole, priorityConsole)

	coreDebug := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerDebug, priorityDebug)
	coreInfo := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerInfo, priorityInfo)
	coreWarn := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerWarn, priorityWarn)
	coreError := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerError, priorityError)
	corePanic := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerPanic, priorityPanic)
	coreFatal := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), writeSyncerFatal, priorityFatal)

	return zap.New(zapcore.NewTee(coreConsole, coreDebug, coreInfo, coreWarn, coreError, corePanic, coreFatal), zap.AddCaller()), nil
}
