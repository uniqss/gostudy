package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 myint
	var i2 = 100
	i1 = 200
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name, s1)

	//i1 = i2 // cannot use i2 (type int) as type myint in assignment

	//name = s1 // cannot use s1 (type string) as type mystr in assignment

	fmt.Printf("%T, %T, %T, %T\n", i1, i2, name, s1)

	fmt.Println("=====================")
	res1 := fun1()
	fmt.Println(res1(10, 20))

	fmt.Println("======================")
	var i3 myint2
	i3 = 1000
	fmt.Println(i3)
	i3 = i2
	fmt.Println(i3)
	fmt.Printf("%T, %T, %T\n", i1, i2, i3)
}

type myint int
type mystr string

type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

//类型别名
type myint2 = int
