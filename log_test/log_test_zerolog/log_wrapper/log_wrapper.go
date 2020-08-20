package log_wrapper

import (
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

func InitLog(serverName string, serverNameMini string) bool {
	fAll, _ := os.OpenFile("./" + serverName + "-all.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fDebug, _ := os.OpenFile("./" + serverName + "-debug.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fInfo, _ := os.OpenFile("./" + serverName + "-info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fWarn, _ := os.OpenFile("./" + serverName + "-warn.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fError, _ := os.OpenFile("./" + serverName + "-error.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fFatal, _ := os.OpenFile("./" + serverName + "-fatal.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

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

	return true
}

