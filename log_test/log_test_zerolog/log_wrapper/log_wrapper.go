package log_wrapper

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

var MainLog zerolog.Logger

type FilteredWriter struct {
	w     zerolog.LevelWriter
	level zerolog.Level
}

func (w *FilteredWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}
func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level == w.level {
		return w.w.WriteLevel(level, p)
	}
	return len(p), nil
}

func InitLog(serverName string, serverNameMini string, logFolder string) error {
	logFileFlag := os.O_APPEND|os.O_CREATE|os.O_WRONLY
	var logFilePerm os.FileMode = 0644
	if logFolder == "" {
		logFolder = "./"
	} else {
		err := os.MkdirAll(logFolder, logFilePerm)
		if err != nil {
			fmt.Println("log_wrapper InitLog os.MkdirAll failed. err:", err)
			return err
		}
	}
	fAll, _ := os.OpenFile(logFolder + serverName + "-all.log", logFileFlag, logFilePerm)
	fDebug, _ := os.OpenFile(logFolder + serverName + "-debug.log", logFileFlag, logFilePerm)
	fInfo, _ := os.OpenFile(logFolder + serverName + "-info.log", logFileFlag, logFilePerm)
	fWarn, _ := os.OpenFile(logFolder + serverName + "-warn.log", logFileFlag, logFilePerm)
	fError, _ := os.OpenFile(logFolder + serverName + "-error.log", logFileFlag, logFilePerm)
	fFatal, _ := os.OpenFile(logFolder + serverName + "-fatal.log", logFileFlag, logFilePerm)

	writerDebug := zerolog.MultiLevelWriter(fDebug)
	writerInfo := zerolog.MultiLevelWriter(fInfo)
	writerWarn := zerolog.MultiLevelWriter(fWarn)
	writerError := zerolog.MultiLevelWriter(fError)
	writerFatal := zerolog.MultiLevelWriter(fFatal)

	filteredWriterDebug := &FilteredWriter{writerDebug, zerolog.DebugLevel}
	filteredWriteInfo := &FilteredWriter{writerInfo, zerolog.InfoLevel}
	filteredWriterWarn := &FilteredWriter{writerWarn, zerolog.WarnLevel}
	filteredWriterError := &FilteredWriter{writerError, zerolog.ErrorLevel}
	filteredWriterFatal := &FilteredWriter{writerFatal, zerolog.FatalLevel}

	w := zerolog.MultiLevelWriter(fAll, zerolog.ConsoleWriter{Out: os.Stdout},
		filteredWriterDebug, filteredWriteInfo, filteredWriterWarn, filteredWriterError, filteredWriterFatal)

	MainLog = zerolog.New(w).With().Str("stype", serverNameMini).Timestamp().Logger()

	return nil
}

