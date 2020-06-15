package main

import "fmt"

func main() {
	var num1 int
	num1 = 30
	fmt.Println("num1的数值是：", num1)

	var num2 int = 15
	fmt.Println("num2的数值是：", num2)

	var name = "王二狗"
	fmt.Printf("name 的数值是：%s\n", name)

	sum := 100
	fmt.Printf("sum 的数值是：%d\n", sum)

	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	var n1, f1, s1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, s1)

	var (
		studentName = "李小花"
		age         = 18
		gender      = "女"
	)
	fmt.Printf("学生姓名：%s 年龄：%d 性别：%s\n", studentName, age, gender)
}
