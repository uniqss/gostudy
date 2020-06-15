package main

import (
	"fmt"
)

func main() {
	a := 10
	a += 5
	fmt.Println("a:", a)

	b := [4]int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		fmt.Printf("%d\t", b[i])
	}

	fmt.Println()

	fmt.Printf("%T\n", fun1)
	fmt.Println(fun1)

	var c func(int, int)
	fmt.Println(c)

	c = fun1
	fmt.Println(c)

	fun1(10, 20)
	c(100, 200)

	res1 := fun2
	res2 := fun2(1, 2)
	fmt.Printf("%T, %v\n", res1, res1)
	fmt.Printf("%T, %v\n", res2, res2)
}
func fun2(a, b int) int{
	return a + b
}
func fun1(a, b int) {
	fmt.Println(a, b)
}
