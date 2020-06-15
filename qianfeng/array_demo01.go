package main

import "fmt"

func main() {
	/*
		数据类型
		基本类型：整数、浮点、布尔、字符串
		复合类型：array,slice, map, struct, pointer, function, channel
	*/

	var num1 int
	num1 = 100
	num1 = 200
	fmt.Println(num1)
	fmt.Printf("%p\n", &num1)

	var arr1 [4]int
	fmt.Printf("%p\n", &arr1)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0])
	fmt.Println(arr1[2])
	fmt.Print(arr1)
	//fmt.Println(arr1[4]) // invalid array index 4 (out of bounds for 4-element array)

	fmt.Println("数组的长度：", len(arr1)) // 实际存储的数据量
	fmt.Println("数组的容量：", cap(arr1)) // 最大能存储的数据量

	var a [4]int
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c = [5]int{1, 2, 3}
	fmt.Println(c)

	var d = [5]int{1: 1, 3: 2}
	fmt.Println(d)

	var e = [5]string{"rose", "王二狗", "ruby"}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(f)
	fmt.Println(len(f))
	g := [...]int{1: 3, 6: 5}
	fmt.Println(g)
	fmt.Println(len(g))
}
