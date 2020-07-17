package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person:", p.name)
}

func (p People) show2() {
	fmt.Println("People:", p.name)
}

type People = Person
type Student struct {
	Person
	People
}

func main() {
	var s Student
	//s.name = "王二狗" // ambiguous selector s.name
	s.Person.name = "王二狗"
	s.Person.show()

	fmt.Printf("%T, %T\n", s.Person, s.People)

	s.People.name = "李小花"
	s.People.show()
	s.People.show2()
	s.Person.show()
	s.Person.show2()
}
