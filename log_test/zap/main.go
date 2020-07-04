package main

import "go.uber.org/zap"

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()
	log.Info("this is log info")
	log.Debug("this is debug log")
	log.Warn("this is warn log")
	log.Error("this is error log")
	//log.Fatal("this is fatal log")
	log.Info("this is info log")

	suger := log.Sugar()
	suger.Infow("hello world encountered a config error.",
		"config", "asdf.xml",
		"reason", "key not found",
		"key", "worldhello",
		"hello", "world",
		"world", 1234)
}
