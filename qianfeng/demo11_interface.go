package main

import "fmt"

func main() {
	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("================")
	var a1 A = cat
	a1.test1()

	fmt.Println("==================")
	var b1 B = cat
	b1.test2()

	fmt.Println("==================")
	var c1 C = cat
	c1.test1()
	c1.test2()
	c1.test3()
}

type A interface {
	test1()
}
type B interface {
	test2()
}
type C interface {
	A
	B
	test3()
}
type Cat struct {

}
func (c Cat) test1() {
	fmt.Println("test1()...")
}
func (c Cat) test2() {
	fmt.Println("test2()...")
}
func (c Cat) test3() {
	fmt.Println("test3()...")
}