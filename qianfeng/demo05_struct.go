package main

import "fmt"

func main() {
	p1 := Person{name:"张三", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	s1 := Student{Person{"李四", 17}, "武汉理工"}
	fmt.Println(s1)

	s2 := Student{Person:Person{name:"rose", age:19}, school: "北京大学"}
	fmt.Println(s2)

	var s3 Student
	s3.Person.name = "王五"
	s3.Person.age = 19
	s3.school = "清华大学"
	fmt.Println(s3)

	s3.name = "ruby"
	s3.age = 16
	fmt.Println(s3)

	fmt.Println(s1.name, s1.age, s1.school)
	fmt.Println(s2.name, s2.age, s2.school)
	fmt.Println(s3.name, s3.age, s3.school)

	var s4 SuperMan
	s4.Person.name = "adsf"
	s4.Person.age = 1234
	s4.name = "SuperMan1"
	fmt.Println(s4)
}

type Person struct {
	name string
	age int
}

type Student struct {
	Person // 模拟继承
	school string // 子类的新增属性
}

type SuperMan struct {
	Person
	name string
}
