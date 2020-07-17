package main

import "fmt"

func main() {
	var a int8
	a = 10

	var b int16
	b = int16(a)

	fmt.Println(a, b)

	f1 := 4.13
	var c int
	c = int(f1)
	fmt.Println(f1, c)

	f1 = float64(a)
	fmt.Println(f1, a)

	//b1 := true
	//a = int8(b1)// cannot convert b1 (type bool) to type int8

	sum := f1 + 100
	fmt.Printf("%T, %f\n", sum, sum)
}
