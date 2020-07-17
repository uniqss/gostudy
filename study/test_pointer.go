package main

import "fmt"

func main() {
	var x int
	x = 1
	var p *int
	p = &x

	var y int = 100

	fmt.Println("x:", x)
	*p = *p + 1
	fmt.Println("x:", x)

	fmt.Println("x:", x)
	add(x)
	fmt.Println("x:", x)

	fmt.Println("x:", x)
	add_pointer(&x)
	fmt.Println("x:", x)

	swap(&x, &y)
	fmt.Println(x, y)
}

func add(q int) {
	q = q + 1
}

func add_pointer(r *int) {
	*r = *r + 1
}

func swap(x *int, y *int) {
	var temp = *x
	*x = *y
	*y = temp
}
