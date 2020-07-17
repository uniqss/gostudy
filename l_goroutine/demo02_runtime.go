package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("os:", runtime.GOOS)

	fmt.Println("NumCPU:", runtime.NumCPU())

	n := runtime.GOMAXPROCS(8)
	fmt.Println("GOMAXPROCS:", n)
	//
	//go func(){
	//	for i:=0;i <5;i++{
	//		fmt.Println("goroutine...")
	//	}
	//}()
	//
	//for i:=0;i < 4;i++ {
	//	runtime.Gosched()
	//	fmt.Println("main...")
	//}

	go func() {
		fmt.Println("goroutine start..")
		fun()
		fmt.Println("goroutine stop...")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("main over..")
}

func fun() {
	defer fmt.Println("fun dever...")
	//return
	runtime.Goexit()
	fmt.Println("fun函数...")
}
