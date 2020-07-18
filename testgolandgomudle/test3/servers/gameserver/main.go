package main

import (
	"github.com/cihub/seelog"
	"library_wrappers/log_wrapper"
)

func main() {
	seelog.Info("main seelog")
	log_wrapper.TestLog()
}
