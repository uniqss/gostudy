package main

import (
	"fmt"
	"strings"
)

var (
	ch = make(chan int)
)

func for_select_default() {
	for {
		select {
		case <-ch:
			fmt.Println("ch channel triggered")
			// this default should be deleted!!!
		default:
		}
	}
}

func main() {
	for i:=0;i < 10;i++ {
		go for_select_default()
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
