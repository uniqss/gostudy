package main

import "fmt"

func main() {
	/*
	fmt.Println("main begin")
	defer test1("defer 1")
	fmt.Println("王二狗")
	defer fmt.Println("defer 2")
	fmt.Println("main end")
	 */

	/*
	a := 2
	fmt.Println(a)
	defer fun2(a)
	a++
	fmt.Println("main中：", a)
	 */

	fmt.Println(fun3())
}

func fun3()int{
	fmt.Println("fun3")
	defer test1("haha")
	return 0
}

func fun2(a int){
	fmt.Println("fun2中：", a)
}

func test1(s string) {
	fmt.Println(s)
}
