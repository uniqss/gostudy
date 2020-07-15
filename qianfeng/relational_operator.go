package main

import "fmt"

func main() {
	/*
		> < >= <= == !=
	*/
	a := 3
	b := 5
	c := 3
	res1 := a > b
	res2 := b > c

	fmt.Printf("%T, %t\n", res1, res1)
	fmt.Printf("%T, %t\n", res2, res2)

	res3 := a == b
	res4 := a == c
	fmt.Printf("%T, %t\n", res3)
	fmt.Println(res4)

	fmt.Println(a != b, a != c)
}
