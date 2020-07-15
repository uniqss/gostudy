package main

import "fmt"

func main() {
	var a chan int
	fmt.Printf("%T, %v\n", a, a)

	if a == nil {
		fmt.Println("chan是nil,不能使用，创建之")
		a = make(chan int)
		fmt.Println(a)
	}

	test1(a)
}

func test1(ch chan int){
	fmt.Printf("%T, %v\n", ch, ch)
}