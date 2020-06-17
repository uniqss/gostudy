package main

import "fmt"

func main() {
	p1 := Person{name: "王二狗", age: 30}
	fmt.Println(p1.name, p1.age)
	p1.eat()

	s1 := Student{Animal{18}, Person{"Ruby", 18}, "武汉理工"}
	fmt.Println(s1.name)
	fmt.Println(s1.Person.age)
	fmt.Println(s1.school)

	s1.eat()
	s1.study()
}

type Animal struct {
	age int
}

type Person struct {
	name string
	age  int
}

type Student struct {
	Animal
	Person
	school string
}

func (a Animal) eat() {
	fmt.Println("动物的方法 吃吃吃，啃啃啃 ")
}

func (p Person) eat() {
	fmt.Println("父类的方法 吃饭，吃窝窝头")
}

func (s Student) study(){
	fmt.Println("子类新增的方法，学习了")
}

func (s Student) eat() {
	fmt.Println("子类新增的方法，学生吃食堂")
}
