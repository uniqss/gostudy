package log_wrapper

func init() {
	var logDir string
	var logLevel string
	var logToConsole bool
	logDir = "./log/"
	logLevel = "debug"
	logToConsole = true

	InitLogger(logDir, logLevel, logToConsole)
}
