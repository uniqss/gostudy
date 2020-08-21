package main
//
//import (
//	"fmt"
//	"go.uber.org/zap"
//	"strings"
//)
//
//func NewLogger2() *zap.Logger {
//	cfg := zap.NewProductionConfig()
//	cfg.OutputPaths = []string{
//		"./log/zaplog.log",
//	}
//	log, err := cfg.Build()
//	if err != nil {
//		fmt.Println("err:", err)
//	}
//
//	defer log.Sync()
//	return log
//}
//
//func main() {
//	log := NewLogger2()
//	log.Debug("this is debug log中文")
//	log.Info("This is info log 哈哈")
//	log.Warn("This is warn log 哈哈")
//	log.Error("This is Error log 哈哈")
//	//log.Fatal("This is Fatal log 哈哈")
//
//	var input string
//	for {
//		fmt.Scanln(&input)
//		input = strings.ToLower(input)
//		if input == "e" || input == "exit" {
//			break
//		} else {
//			log.Info("input", zap.String("input", input))
//		}
//	}
//}
