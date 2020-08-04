package main

import "fmt"

type Something struct {
	someA string
	someB int
}

func NewSomething() *Something {
	return &Something{someA: "", someB: 0}
}

func someFunc1(someInterface interface{}) {
	s := someInterface.(*Something)
	s.someA = "asdf"
	s.someB = 1234
}

func (s *Something) print(){
	fmt.Println("Something: someA:", s.someA, " someB:", s.someB)
}

func main() {
	s := NewSomething()
	s.print()
	someFunc1(s)
	// this means    interface{}    is the same as    void*            !!!!!!!
	s.print()
}
