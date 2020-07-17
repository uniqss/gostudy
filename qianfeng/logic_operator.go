package main

import "fmt"

func main() {
	/*
		&& ||
		操作数必须是bool，运算结果也是bool
	*/

	f1 := true
	f2 := false
	f3 := true
	res1 := f1 && f2
	fmt.Println("res1:", res1)

	res2 := f1 && f2 && f3
	fmt.Println("res2:", res2)

	res3 := f1 || f2
	fmt.Println("res3:", res3)

	res4 := f1 || f2 || f3
	fmt.Println("res4:", res4)
	fmt.Println(false || false || false)

	fmt.Printf("f1:%t, !f1:%t\n", f1, !f1)
	fmt.Printf("f2:%t !f2:%t\n", f2, !f2)

	a := 3
	b := 2
	c := 5
	res5 := a > b && c%a == b && a < (c/b)
	fmt.Println(res5)

	res6 := b*2 < c || a/b != 0 || c/a > b
	fmt.Println(res6)
	res7 := !(c/a == b)
	fmt.Println(res7)
}
