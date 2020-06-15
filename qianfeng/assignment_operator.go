package main

import "fmt"

func main() {
	/*
		= += -= *= /= %= <<= >>= &= |= ^= ...
	*/
	var a int
	a = 3
	fmt.Println(a)

	a += 4
	fmt.Println(a)

	a -= 3
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 3
	fmt.Println(a)
	a %= 1
	fmt.Println(a)
}
