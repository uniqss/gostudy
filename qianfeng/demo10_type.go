package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%T\n", a)

	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", b)

	c := []int{1, 2, 3, 4}
	fmt.Printf("%T\n", c)

	d := make(map[int]string)
	fmt.Printf("%T\n", d)

	fmt.Printf("%T\n", fun1)
	fmt.Printf("%T\n", fun2)
	fmt.Printf("%T\n", fun3)
	fmt.Printf("%T\n", fun4)
}

func fun1() {}
func fun2(a int) int {
	return 0
}
func fun3(a float64, b, c int) (int, int) {
	return 0, 0
}
func fun4(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}
