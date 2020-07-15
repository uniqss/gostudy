package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go sendData(ch1)

	for v := range ch1 {
		fmt.Println("读的数据：", v)
	}
	fmt.Println("main over...")
}

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch1 <- i
	}
	close(ch1)
}
