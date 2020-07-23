package main

import (
	"fmt"
	"runtime"
	"strings"
)

func ping(){
	for {
		// every program should have at most one runtime.Gosched() call !!!
		runtime.Gosched()
	}
}

func pong(){
	for  {
		// every program should have at most one runtime.Gosched() call !!!
		runtime.Gosched()
	}
}

func main()  {

	// every program should have at most one runtime.Gosched() call !!!
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

