package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
			fmt
				输出
					Print()
					Printf()
					Println()
		%v
		%T
		%t
		%s
		%f
		%d
		%b
		%o
		%x %X
		%c
		%p

	*/
	fmt.Println("main")
	a := 100
	b := 3.14
	c := true
	d := "hello world"
	e := "Ruby"
	f := 'A'
	fmt.Printf("%T, %b\n", a, a)
	fmt.Printf("%T, %f\n", b, b)
	fmt.Printf("%T, %t\n", c, c)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %s\n", e, e)
	fmt.Printf("%T, %d, %c\n", f, f, f)
	fmt.Println("-------------------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	/*
		scanln
	*/
	/*
		var x int
		var y float64
		fmt.Println("请输入一个整数，一个浮点数：")
		fmt.Scanln(&x, &y)
		fmt.Printf("a的数值:%d， b的数值：%f\n", x, y)

		fmt.Scanf("%d, %f", &x, &y)
		fmt.Printf("x:%d, y:%f\n", x, y)
	*/

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)
}
