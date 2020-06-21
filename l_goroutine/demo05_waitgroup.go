package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go fun1()
	go fun2()

	fmt.Println("main wg.Wait before")
	wg.Wait()
	fmt.Println("main wg.Wait after")
}

func fun1() {
	for i := 1; i < 10; i++ {
		fmt.Println("fun1()函数中打印。..111", i)
	}

	wg.Done()
}

func fun2() {
	defer wg.Done()

	for i := 1; i < 10; i++ {
		fmt.Println("fun2()函数中打印。。222", i)
	}
}
