package main

import "fmt"

func main() {
	p1 := Person{"王二狗", 30, "男", "北京市"}
	fmt.Println(p1)
	fmt.Printf("%p, %T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p, %T\n", &p2, p2)

	p2.name = "李小花"
	fmt.Println(p2)
	fmt.Println(p1)

	// struct pointer
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p, %T\n", pp1, pp1)

	pp1.name = "王五"
	fmt.Println(pp1)
	fmt.Println(p1)

	// new
	pp2 := new(Person)
	fmt.Printf("%T\n", pp2)
	fmt.Println(pp2)
	pp2.name = "Jerry"
	pp2.age = 20
	pp2.gender = "男"
	pp2.address = "广州"
	fmt.Println(pp2)

	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
}

type Person struct {
	name string
	age int
	gender string
	address string
}