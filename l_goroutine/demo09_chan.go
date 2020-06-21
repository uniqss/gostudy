package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i:", i)
		}
		// 循环结束就往通道中写数据，表结束
		ch1 <- true
		fmt.Println("结束")
	}()

	data := <-ch1
	fmt.Println(data)

	fmt.Printf("main over..")

}
