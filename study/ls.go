package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "the dir to list")
	flag.Parse()

	//f, err := os.Open(".")
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
	}
	f.Close()
}
