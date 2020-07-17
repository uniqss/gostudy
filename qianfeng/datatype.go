package main

import "fmt"

func main() {
	fmt.Println("func main")
	/*
		GO语言数据类型：
		1、基本数据类型：
			布尔类型：bool
				取值：true, false
			数值类型：
				整数：
				浮点： 生活中的小数
				复数：complex
			字符串：string
		2、复合数据类型
			array, slice, map, function, pointer, struct, struct, interface, channel...
	*/

	//1.bool
	var b1 bool
	b1 = true
	fmt.Printf("%T, %t\n", b1, b1)
	b2 := false
	fmt.Printf("%T, %t\n", b2, b2)

	// 2.整数
	var i1 int8
	i1 = 100
	fmt.Println(i1)
	var i2 uint8
	i2 = 200
	fmt.Println(i2)

	var i3 int
	i3 = 1000
	fmt.Println(i3)

	var i4 int64
	//i4 = i3 // cannot use i3 (type int) as type int64 in assignment
	fmt.Println(i4)

	var i5 uint8
	i5 = 100
	var i6 byte
	i6 = i5
	fmt.Println(i5, i6)
	var i7 = 100
	fmt.Printf("%T, %d\n", i7, i7)

	// 3.float
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 4.67
	fmt.Printf("%T, %.2f\n", f1, f1)
	fmt.Printf("%T, %.3f\n", f2, f2)
	fmt.Println(f1)

	var f3 = 2.55
	fmt.Printf("%T\n", f3)
}
