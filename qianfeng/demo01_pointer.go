package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a的数值是：", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("a的地址是：%p\n", &a)

	// 创建一个指针，用来存储a的地址
	var p1 *int
	fmt.Println(p1) // nil
	p1 = &a
	fmt.Println("p1的数值：", p1)
	fmt.Println("p1自己的地址：", &p1)
	fmt.Println("p1的数值，是a的地址，该地址存储的数据：", *p1)

	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a)

	*p1 = 200
	fmt.Println(a)
	var p2 **int
	fmt.Println(p2)

	p2 = &p1
	fmt.Printf("%T %T %T\n", a, p1, p2)
	fmt.Println("p2的数值：", p2)
	fmt.Println("p2自己的地址：", &p2)
	fmt.Println("p2指向的数值：", *p2)
}
