package main

import (
	"fmt"
	ulog "log_test_zerolog/log_wrapper"
)

func main() {
	err := ulog.InitLog("game_server", "gs", "./log/")
	if err != nil {
		fmt.Println(err)
		return
	}

	ulog.MainLog.Debug().Str("hello", "world").Msg(" Some Debug log 我是中文哈哈")
	ulog.MainLog.Info().Strs("xixi", []string{"haha", "hehe"}).Msg(" Some  Info log 呵呵吼吼")
	ulog.MainLog.Warn().Msg(" Some warning 哦吼！！着火了")
	ulog.MainLog.Error().Msg(" Some Error log 卧槽，，服务器挂了")
	ulog.MainLog.Fatal().Msg(" Some Fatal log! 不的了了。。。机房冒烟了")
}
