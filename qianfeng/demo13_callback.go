package main

import "fmt"

func main() {
	fmt.Printf("%T %p\n", add, add)
	fmt.Printf("%T %p\n", oper, oper)

	fmt.Println(oper(1, 2, add))
	fmt.Println(oper(1, 2, sub))

	fun1 := func(a, b int) int{
		return a * b
	}
	fmt.Println(oper(3, 5, fun1))

	fmt.Println(oper(100, 20, func(a, b int) int{
		if b == 0 {
			fmt.Println("divided by 0")
			return 0
		}
		return a / b
	}))
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func oper(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}
