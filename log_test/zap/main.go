package main

import (
	"fmt"
	"strings"
	"zap/common"
)

func init() {
	common.ZLog = common.NewLogger("./log/worldserver.log", true)
}

func main() {
	zLog := common.ZLog
	zLog.Debug("this is debug log中文")
	zLog.Info("This is info log 哈哈")
	zLog.Warn("This is warn log 哈哈")
	zLog.Error("This is Error log 哈哈")
	//zLog.Fatal("This is Fatal log 哈哈")

	SomeOther()

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		} else {
			zLog.Info("input is:" + input)
		}
	}
}
