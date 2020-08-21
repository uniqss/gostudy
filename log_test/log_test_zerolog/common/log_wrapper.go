package common

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

var ULogger zerolog.Logger

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

func openLogFile(fileName string, flag int, fileMode os.FileMode) (*os.File, error) {
	fDebug, err := os.OpenFile(fileName, flag, fileMode)
	if err != nil {
		fmt.Println("openLogFile failed. fileName:", fileName, " err:", err)
		return nil, err
	}
	return fDebug, nil
}

func InitLog(logFileName string, logFolder string, allLog bool) error {
	logFileFlag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	var logFilePerm os.FileMode = 0644
	if logFolder == "" {
		logFolder = "./"
	} else {
		if !strings.HasSuffix(logFolder, "/") {
			logFolder += "/"
		}
		err := os.MkdirAll(logFolder, logFilePerm)
		if err != nil {
			fmt.Println("log_wrapper InitLog os.MkdirAll failed. err:", err)
			return err
		}
	}

	fDebug, err := openLogFile(logFolder+logFileName+"_debug.log", logFileFlag, logFilePerm)
	if err != nil {
		return err
	}
	fInfo, err := openLogFile(logFolder+logFileName+"_info.log", logFileFlag, logFilePerm)
	if err != nil {
		return err
	}
	fWarn, err := openLogFile(logFolder+logFileName+"_warn.log", logFileFlag, logFilePerm)
	if err != nil {
		return err
	}
	fError, err := openLogFile(logFolder+logFileName+"_error.log", logFileFlag, logFilePerm)
	if err != nil {
		return err
	}
	fFatal, err := openLogFile(logFolder+logFileName+"_fatal.log", logFileFlag, logFilePerm)
	if err != nil {
		return err
	}

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

	if allLog {
		fAll, err := openLogFile(logFolder+logFileName+"_all.log", logFileFlag, logFilePerm)
		if err != nil {
			return err
		}
		logLevelWriter := zerolog.MultiLevelWriter(
			fAll,
			zerolog.ConsoleWriter{Out: os.Stdout},
			filteredWriterDebug, filteredWriteInfo, filteredWriterWarn, filteredWriterError, filteredWriterFatal)

		ULogger = zerolog.New(logLevelWriter)
	} else {
		logLevelWriter := zerolog.MultiLevelWriter(
			zerolog.ConsoleWriter{Out: os.Stdout},
			filteredWriterDebug, filteredWriteInfo, filteredWriterWarn, filteredWriterError, filteredWriterFatal)

		ULogger = zerolog.New(logLevelWriter)
	}

	// usage
	//uLog := common.ULogger.With().Timestamp().Str("func", "main").Logger()
	//uLog.Info().Str("input", input).Msg("user input")

	return nil
}
