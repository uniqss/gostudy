package main

import "fmt"

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan<- int)
	//ch3 := make(<-chan int)

	//ch1 <- 100
	//data := <-ch1

	//ch2 <- 1000 // invalid operation: <-ch2 (receive from send-only type chan<- int)
	//data := <-ch2

	//data := <-ch3
	//ch3 <- 2000 // invalid operation: ch3 <- 2000 (send to receive-only type <-chan int)

	go fun1(ch1)
	data := <-ch1

	//go fun1(ch2)
	//data := <- ch2 // invalid operation: <-ch2 (receive from send-only type chan<- int)
	fmt.Println("main data:", data)

	go fun2(ch1)
	ch1 <- 200
	fmt.Println("main over...")
}

func fun1(ch chan<- int) {
	ch <- 100
	fmt.Println("fun1函数结束。。。")
}

func fun2(ch <-chan int) {
	data := <-ch
	fmt.Println("fun2 data:", data)
}
