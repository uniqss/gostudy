package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	flag.IntVar(&port, "p", 60000, "listen port. default 8080")
	flag.StringVar(&dir, "dir", "./", "local dir to share")

	flag.Parse()

	h := http.FileServer(http.Dir(dir))
	//err := http.ListenAndServe(":8080", h)
	fmt.Println("host:", host, " port:", port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), h)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Printf("terminating...")
}
