package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
start_share.bat:
start http_curr_dir.exe -dir ./share
*/

func main() {
	var host string
	var port int
	var dir string

	flag.StringVar(&host, "h", "0.0.0.0", "listen host. default 0.0.0.0")
	flag.IntVar(&port, "p", 60000, "listen port. default 60000")
	flag.StringVar(&dir, "dir", "./share", "local dir to share")

	flag.Parse()

	state, err := os.Stat(dir)
	if err != nil && dir != "./" {
		if state != nil && !state.IsDir() {
			fmt.Println("the dir is a file not a folder. using ./ as default. dir:", dir)
			dir = "./"
		} else {
			fmt.Println("dir not exists. trying to create. dir:", dir)
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Println("create failed. using ./ as default")
				dir = "./"
			}
		}
	}

	h := http.FileServer(http.Dir(dir))
	//err := http.ListenAndServe(":8080", h)
	fmt.Println("host:", host, " port:", port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), h)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Printf("terminating...")
}
