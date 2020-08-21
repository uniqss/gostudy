package main

import (
	"fmt"
	"log_test_zerolog/common"
	"strings"
)

func testLog(){
	uLog := common.ULogger.With().Timestamp().Str("func", "testLog").Logger()

	uLog.Debug().Str("hello", "world").Msg(" Some Debug log 我是中文哈哈")
	uLog.Info().Strs("xixi", []string{"haha", "hehe"}).Msg(" Some  Info log 呵呵吼吼")
	uLog.Warn().Msg(" Some warning 哦吼！！着火了")
	uLog.Error().Msg(" Some Error log 卧槽，，服务器挂了")
	// !!! fatal log will exit the program
	//uLog.MainLog.Fatal().Msg(" Some Fatal log! 不的了了。。。机房冒烟了")

}

func main() {
	err := common.InitLog("gameserver", "./log", true)
	if err != nil {
		fmt.Println(err)
		return
	}

	testLog()


	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		} else {
			uLog := common.ULogger.With().Timestamp().Str("func", "main").Logger()
			uLog.Info().Str("input", input).Msg("user input")
		}
	}
}
