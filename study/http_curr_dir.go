package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var host string
	var port int

	flag.StringVar(&host, "h", "0.0.0.0", "listen host. default 0.0.0.0")
	flag.IntVar(&port, "p", 8080, "listen port. default 8080")

	flag.Parse()

	h := http.FileServer(http.Dir("./"))
	//err := http.ListenAndServe(":8080", h)
	fmt.Println("host:", host, " port:", port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), h)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Printf("terminating...")
}
