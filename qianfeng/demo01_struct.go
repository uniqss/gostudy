package main

import "fmt"

func main() {
	// 初始化方法1
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 30
	p1.gender = "男"
	p1.address = "北京市"
	fmt.Printf("姓名：%s年龄:%d性别：%s地址%s\n", p1.name, p1.age, p1.gender, p1.address)

	p2 := Person{}
	p2.name = "Ruby"
	p2.age = 28
	p2.gender = "女"
	p2.address = "上海市"
	fmt.Printf("姓名：%s年龄:%d性别：%s地址%s\n", p2.name, p2.age, p2.gender, p2.address)

	p3 :=Person{name:"如花",age:20,gender: "女", address: "杭州"}
	fmt.Println(p3)
	p4 := Person{
		name:"隔壁老王",
		age:40,
		gender: "男",
		address: "武汉市",
	}
	fmt.Println(p4)

	p5 := Person{"李小花", 25, "女", "成都"}
	fmt.Println(p5)

}

type Person struct {
	name string
	age int
	gender string
	address string
}