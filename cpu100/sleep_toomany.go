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
