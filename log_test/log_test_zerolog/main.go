package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
)

func main() {
	tmpfile, err := ioutil.TempFile("", "testtempfile")
	if err != nil {
		fmt.Println("ioutil.TempFile failed. err:", err)
		return
	}

	log := zerolog.New(tmpfile)
	log.Info().Str("log key1", "log value str 1").
		Str("log key2", "log value str 2").Msg("sample log message")
}
