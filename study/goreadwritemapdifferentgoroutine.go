package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var gmap = make(map[int32]string)
var working = true

func read() {
	for working {
		key := rand.Int31() % 100
		print, ok := gmap[key]
		if ok && print != "" {
			fmt.Println(print)
		} else {
			fmt.Println(strconv.Itoa(int(key)) + " not found")
		}
	}
}

func write() {
	for working {
		key := rand.Int31() % 100
		value := rand.Int31()
		str := strconv.Itoa(int(value))
		gmap[key] = str
	}
}

func main() {
	go read()
	go write()

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			working = false
			break
		}
	}
	time.Sleep(time.Second * 5)
}
