package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		非缓冲通道：make(chan T)
		缓冲通道：make(chan T, capacity)
	*/
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1)) // 0 0
	//ch1 <- 100
	//data := <-ch1
	//fmt.Println(data)

	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 100
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 200
	ch2 <- 300
	ch2 <- 400
	ch2 <- 500
	fmt.Println(len(ch2), cap(ch2))

	// all goroutines are asleep - deadlock!
	//ch2 <- 600

	fmt.Println("==============================")
	ch3 := make(chan string, 4)
	go sendData(ch3)

	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("读完了。。", ok)
			break
		}
		fmt.Println("\t读取的数据是:", v)
	}
	fmt.Printf("main over..")
}

func sendData(ch chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("子goroutine中写出第%d个数\n", i)
		ch <- "数据" + strconv.Itoa(i)
	}
	close(ch)
}
