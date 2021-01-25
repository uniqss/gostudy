package main

import (
	"time"

	"github.com/akath19/gin-zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	//Gin Router
	router := gin.New()
	//Zap logger
	zap, _ := zap.NewProduction()
	//Add middleware to Gin, requires sync duration & zap pointer
	router.Use(ginzap.Logger(3 * time.Second, zap))
	//Other gin configs
	router.Use(gin.Recovery())

	//Logger will work outside Gin as well
	zap.Warn("Warning")
	zap.Error("Error")


	// this is a total package of shit
	zap.Info("Info", zap.String("aaa", "bbb"))
	zap.Info("Info", "asdf")

	//Start Gin
	router.Run(":8080")
}