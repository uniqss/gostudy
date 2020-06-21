package main

import "fmt"

func main() {
	/*
		chan <- data 发送，写
		data <- chan 获取，读
	*/
	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)

	data := <-ch1
	fmt.Println("子goroutine传来：", data)

	ch1 <- "我是main"
	<-done
	fmt.Printf(" main over...")
}

func sendData(ch1 chan string, done chan bool) {
	ch1 <- "我是韩茹"

	data := <-ch1
	fmt.Println("sendData, data:", data)
	done <- true
}
