package common

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
)

var ZLog *zap.Logger

func NewLogger(logFileName string) *zap.Logger {
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

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		logFileName,
	}
	log, err := cfg.Build()
	if err != nil {
		fmt.Println("NewLogger cfg.Build zap log got some error. err:", err)
	}

	defer log.Sync()
	return log
}
