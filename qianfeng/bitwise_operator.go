package main

import "fmt"

func main() {
	/*
		位运算符，按位计算
		&与
		|或
		^异或
		&^位清空
	*/
	a := 60
	b := 13
	fmt.Printf("a:%d, %b\n", a, a)
	fmt.Printf("b:%d, %b\n", b, b)

	res1 := a & b
	fmt.Println(res1)

	res2 := a | b
	fmt.Println(res2)

	res3 := a ^ b
	fmt.Println(res3)

	res4 := a &^ b
	fmt.Println(res4)

	res5 := ^a
	fmt.Println(res5)

	c := 8
	res6 := c << 2
	fmt.Println(res6)

	res7 := c >> 2
	fmt.Println(res7)
}
