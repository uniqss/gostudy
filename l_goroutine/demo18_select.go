package main

import (
	"fmt"
	"time"
)

func main() {
	// 分支语句：if switch select
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中获取的数据：", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中读取的数据：", num2)
		} else {
			fmt.Println("ch2已关闭")
		}
	default:
		fmt.Println("default")
	}
}
