package main
//
//import (
//	"flag"
//	"fmt"
//	"github.com/rs/zerolog"
//	"github.com/rs/zerolog/log"
//	"io/ioutil"
//)
//
//func main() {
//	test := 1
//	if test == 1 {
//		tmpfile, err := ioutil.TempFile("", "testtempfile")
//		if err != nil {
//			fmt.Println("ioutil.TempFile failed. err:", err)
//			return
//		}
//
//		log := zerolog.New(tmpfile)
//		log.Info().Str("log key1", "log value str 1").
//			Str("log key2", "log value str 2").Msg("sample log message")
//	} else if test == 2 {
//		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
//		debug := flag.Bool("debug", false, "sets log level to debug")
//
//		flag.Parse()
//
//		// Default level for this example is info, unless debug flag is present
//		zerolog.SetGlobalLevel(zerolog.InfoLevel)
//		if *debug {
//			zerolog.SetGlobalLevel(zerolog.DebugLevel)
//		}
//
//		log.Debug().Msg("This message appears only when log level set to Debug")
//		log.Info().Msg("This message appears when log level set to Debug or Info")
//
//		if e := log.Debug(); e.Enabled() {
//			// Compute log output only if enabled.
//			value := "bar"
//			e.Str("foo", value).Msg("some debug message")
//		}
//	}
//}
