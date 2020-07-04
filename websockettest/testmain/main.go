package main

import (
	log "github.com/sirupsen/logrus"
	"time"
	"websockettest/common"
)

func main() {
	log.Info("server startup...")

	service := common.NewService()

	mp := common.NewTestMsgProc()
	service.Register(mp)

	err := service.Init("0.0.0.0:8080")
	if err != nil {
		log.Error("service.init failed. err:", err)
		return
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)
	}

	log.Info("server exit...")
}
