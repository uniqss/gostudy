package main

import "fmt"

func main() {
	var a func(a int)
	a = fun1
	a(123)

	arr1 := fun2()
	fmt.Printf("arr1 %T, %p, %v\n", arr1, &arr1, arr1)

	arr2 := fun3()
	fmt.Printf("%p\n", arr2)
	fmt.Printf("arr2 %T, %p, %v\n", arr2, &arr2, arr2)
}

func fun3()*[4]int{
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("%p\n", arr)
	return &arr
}

func fun2()[4]int{
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun1(a int){
	fmt.Println("fun1()... a:", a)
}