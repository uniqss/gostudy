package main

import "fmt"

func main() {
	res1 := add(1, 2)
	fmt.Println(res1)

	res2 := sub(5, 3)
	fmt.Println(res2)

	res3 := oper(5, 6, add)
	fmt.Println(res3)

	func1 := func(a, b int) int {
		return a * b
	}

	res4 := oper(5, 6, func1)
	fmt.Println(res4)

	res5 := oper(100, 8, func(a, b int) int {
		if b == 0 {
			fmt.Println("divided by 0")
			return 0
		}
		return a / b
	})

	fmt.Println(res5)
}

func oper(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
