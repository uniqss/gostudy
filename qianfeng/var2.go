package main

import "fmt"

var a = 10000
var b int = 2000

//c := 3000 // syntax error: non-declaration statement outside function body

func main() {
	var num int
	num = 100
	fmt.Printf("num的数值是：%d，:%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d，:%p\n", num, &num)

	//fmt.Println(num2) // undefined: num2

	var name string
	//name = 100
	//fmt.Println(name) // cannot use 100 (type untyped int) as type string in assignment

	name = "张二狗"
	fmt.Println(name)

	//var name string = "李小花" // name redeclared in this block
	//fmt.Println(name)

	num, name, gender := 1000, "李小花", "女"
	fmt.Println(num, name, gender)

	fmt.Println("------------------------")
	var m int
	fmt.Println(m)
	var n float64
	fmt.Println(n)
	var s string
	fmt.Println(s)
	var s2 []int
	fmt.Println(s2)
	fmt.Println(s2 == nil)

	//var sum = 100 // sum declared but not used
}
