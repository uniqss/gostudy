package main

import (
	"fmt"
)

func main() {
	go printNum()

	for i := 1; i <= 1000; i++ {
		fmt.Printf("主goroutine中打印字母 A %d\n", i)
	}

	//time.Sleep(1 * time.Second)

	fmt.Printf("main over...")
}

func printNum() {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("子goroutine中打印的数字：%d\n", i)
	}
}
