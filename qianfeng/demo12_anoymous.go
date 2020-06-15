package main

import "fmt"

func main() {
	fun1()
	fun2 := fun1
	fun2()

	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	fun3()
	fun3()

	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	res1 := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(res2)
	fmt.Println(res2(100, 200))
	fmt.Println(res2(111, 222))

	// 回调函数 函数作为参数传到另一个函数中，完成时调用
	// 闭包，函数作为参数返回
}

func fun1() {
	fmt.Println("我是fun1()函数。。。")
}
