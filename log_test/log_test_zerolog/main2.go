package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"log_test_zerolog/log_wrapper"
	"strings"
)

func main() {
	err := log_wrapper.InitLog("game_server", "./log/")
	if err != nil {
		fmt.Println(err)
		return
	}

	uLog := zerolog.New(log_wrapper.LogLevelWriter).With().Timestamp().Str("func", "main").Logger()

	uLog.Debug().Str("hello", "world").Msg(" Some Debug log 我是中文哈哈")
	uLog.Info().Strs("xixi", []string{"haha", "hehe"}).Msg(" Some  Info log 呵呵吼吼")
	uLog.Warn().Msg(" Some warning 哦吼！！着火了")
	uLog.Error().Msg(" Some Error log 卧槽，，服务器挂了")
	// !!! fatal log will exit the program
	//uLog.MainLog.Fatal().Msg(" Some Fatal log! 不的了了。。。机房冒烟了")

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		} else {
			uLog.Info().Str("input", input).Msg("main. user input")
		}
	}
}
