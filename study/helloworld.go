package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	var ss []string
	ss = append(ss, "aaa")
	ss = append(ss, "bbb")
	fmt.Println(ss)
	fmt.Println(ss[2:])
	fmt.Println(ss[1:])
	fmt.Println(ss[0:1])
	fmt.Println(ss[1:2])
	fmt.Println(ss[0:2])
	fmt.Println("asdf")

	var timer *time.Timer
	if timer != nil {
		for {
			select {
			case <-timer.C:
				fmt.Println("asdf")
			}
		}
	}

	fmt.Println("main exiting...")

}
