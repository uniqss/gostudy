package main

import "fmt"

func main() {
	/*
		外层函数中调用内层函数，内层函数操作外层函数的局部变量
		并且外层函数把内层函数作为返回值
		这个内层函数和外层函数的局部变量，统称为闭包结构
	*/
	res1 := increment()
	fmt.Printf("%T\n", res1)
	fmt.Printf("%p\n", &res1)
	v1 := res1()
	fmt.Println(v1)
	v2 := res1()
	fmt.Println(v2)

	res2 := increment()
	fmt.Printf("%p\n", &res2)
	fmt.Println(res2())
}

func increment() func() int { // 外层函数
	// 1.定义了一个变量
	i := 0
	// 2.定义了一个匿名函数，给变量自增并返回
	fun := func() int { // 内置函数
		i++
		return i
	}
	// 3.返回该匿名函数
	return fun
}
