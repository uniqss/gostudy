package main

import log "github.com/sirupsen/logrus"

func main() {

	//logger := log.New()
	//
	//logger.Formatter = &log.TextFormatter{
	//	DisableColors:  true,
	//	FullTimestamp:  true,
	//	DisableSorting: true,
	//}
	//
	//logger.Debug("logrus's log debug")
	//logger.Info("logrus's log info")
	//logger.Warn("logrus's log warn")
	//logger.Error("logrus's log error")
	//logger.Fatal("logrus's log fatal")

	log.Debug("logrus's log debug")
	log.Info("logrus's log info 1")
	log.Info("logrus's log info 2")
	log.Warn("logrus's log warn")
	log.Info("logrus's log info 3")
	log.Error("logrus's log error")
	log.Info("logrus's log info 4")
	log.Fatal("logrus's log fatal")
	log.Info("logrus's log info 5")

	log.Info("server terminating...")

	log.Debugf("logrus's log with format. %d", 123)
}
