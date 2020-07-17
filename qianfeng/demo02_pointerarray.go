package main

import "fmt"

func main() {
	/*
		数组指针：首先是一个指针，一个数组的地址
			*[4]Type
		指针数组：首先是一个数组，存储的数据类型是指针
			[4]*Type
			*[5]float64,数组指针，一个存储了5个浮点类型数据的数组的指针
			*[3]string,数组指针，三个字符串的数组的指针
			[3]*string，指针数组，存储了三个字符串的指针地址的数组
			[5]*float64，指针数组，存储了5个float类型指针的数组
			*[5]*float64，指针数组指针
			*[3]*string，指针数组指针
			**[4]string，指针的指针，存储了4个字符串数组的指针的指针
			**[4]*string，指针的指针，存储了4个字符串指针的数组的指针的指针
	*/
	// 创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	fmt.Printf("%p\n", &arr1)

	// 创建一个指针，存储该数组的地址
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", &p1)

	(*p1)[0] = 100
	fmt.Println(arr1)

	// 数组指针
	p1[0] = 200
	fmt.Println(arr1)

	// 指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)

	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)
	b = 1000
	fmt.Println(arr2)
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("%d\t", *arr3[i])
	}
	fmt.Println()
}
