package main

import (
	log "zap2/log_wrapper"
)

func main() {
	log.Debug("debug", 1234, " aaaa:", 891234798)
	log.Info("info", 1234)
	log.Warn("warn", 1234)
	log.Error("error", 1234)
	//log.Panic("panic", 1234)
	//log.Fatal("fatal", 1234)

	for i := 0; i < 100000; i++ {
		log.Debug("debug", 1234)
		log.Debug("debug", 1234, " aaaa:", 321)
	}
}
