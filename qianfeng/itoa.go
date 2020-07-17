package main

import "fmt"

func main() {
	fmt.Println("asdf")

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c)

	const (
		MALE    = iota
		FEMALE  = iota
		UNKNOWN = iota
	)
	fmt.Println(MALE, FEMALE, UNKNOWN)

	const (
		A = iota
		B
		C
		D = "haha"
		E
		F = 100
		G
		H = iota
		I
	)
	const (
		J = iota
	)
	fmt.Println(A, B, C, D, E, F, G, H, I, J)
}
