package main

import (
	"fmt"
	"strings"
	"time"
)

var ttt = 100

func oneGoRoutine() {
	for {
		time.Sleep(time.Millisecond * 100)
		i := 0
		for i := 0; i < 1000; i++ {
			i++
		}
		ttt = i
	}
}

func main() {

	// in goroutine calling time.Sleep is also expensive if there are too many goroutines !!!
	// so call time.Sleep as LESS as possible !!!
	// if there is a select in the for, time.Sleep should not be called, and the select should not have default !!!
	for i := 0; i < 500000; i++ {
		go oneGoRoutine()
	}

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		}
	}

}
