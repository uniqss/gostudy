package main

import (
	"fmt"
	"runtime"
	"strings"
)

func ping(){
	for {
		runtime.Gosched()
	}
}

func pong(){
	for  {
		runtime.Gosched()
	}
}

func main()  {
	go ping()
	go pong()

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		}
	}

}

