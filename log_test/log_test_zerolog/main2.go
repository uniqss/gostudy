package main

import (
	"log_test_zerolog/log_wrapper"
)

func main() {
	log_wrapper.InitLog("game_server", "gs")

	log_wrapper.MainLog.Debug().Str("hello", "world").Msg(" Some Debug log 我是中文哈哈")
	log_wrapper.MainLog.Info().Strs("xixi", []string{"haha", "hehe"}).Msg(" Some  Info log 呵呵吼吼")
	log_wrapper.MainLog.Warn().Msg(" Some warning 哦吼！！着火了")
	log_wrapper.MainLog.Error().Msg(" Some Error log 卧槽，，服务器挂了")
	log_wrapper.MainLog.Fatal().Msg(" Some Fatal log! 不的了了。。。机房冒烟了")
}
