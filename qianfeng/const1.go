package main

import "fmt"

func main() {
	fmt.Println(100)
	fmt.Println("hello")

	const PATH string = "http://www.baidu.com"
	const PI = 3.14
	const NOTUSED = "notused"
	fmt.Println(PATH)
	fmt.Println(PI)

	//PATH = "www.google.com" // cannot assign to PATH

	const C1, C2, C3 = 100, 3.14, "哈哈"
	const (
		MALE    = 0
		FEMALE  = 1
		UNKNOWN = 3
	)

	const (
		a int = 100
		b
		c string = "ruby"
		d
		e
	)
	fmt.Println(a, b)
	fmt.Println(c, d, e)

	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
