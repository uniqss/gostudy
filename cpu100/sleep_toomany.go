package main

import (
	"fmt"
	"strings"
	"time"
)

func sleep_toomany(){
	for  {
		time.Sleep(time.Millisecond)
	}
}

func main() {

	// in goroutine calling time.Sleep is also expensive if there are too many goroutines !!!          
	// so call time.Sleep as LESS as possible !!!
	// if there is a select in the for, time.Sleep should not be called, and the select should not have default !!!
	for i:=0;i < 20000;i++ {
		go sleep_toomany()
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
