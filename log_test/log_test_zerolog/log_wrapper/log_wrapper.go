package log_wrapper

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

var LogLevelWriter zerolog.LevelWriter

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

func InitLog(logFileName string, logFolder string) error {
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
	fAll, _ := os.OpenFile(logFolder +logFileName+ "_all.log", logFileFlag, logFilePerm)
	fDebug, _ := os.OpenFile(logFolder +logFileName+ "_debug.log", logFileFlag, logFilePerm)
	fInfo, _ := os.OpenFile(logFolder +logFileName+ "_info.log", logFileFlag, logFilePerm)
	fWarn, _ := os.OpenFile(logFolder +logFileName+ "_warn.log", logFileFlag, logFilePerm)
	fError, _ := os.OpenFile(logFolder +logFileName+ "_error.log", logFileFlag, logFilePerm)
	fFatal, _ := os.OpenFile(logFolder +logFileName+ "_fatal.log", logFileFlag, logFilePerm)

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

	LogLevelWriter = zerolog.MultiLevelWriter(fAll, zerolog.ConsoleWriter{Out: os.Stdout},
		filteredWriterDebug, filteredWriteInfo, filteredWriterWarn, filteredWriterError, filteredWriterFatal)

	// usage
	//uLog := zerolog.New(log_wrapper.LogLevelWriter).With().Str("func", "main").Timestamp().Logger()
	//uLog.Debug().Str("hello", "world")

	return nil
}

