package main

import "fmt"

func main() {
	// 值类型 int float string bool array
	// 引用类型 slice

	a1 := [4]int{1, 2, 3, 4}
	a2 := a1
	fmt.Println(a1, a2)
	a1[0] = 100
	fmt.Println(a1, a2)
	fmt.Printf("%p %p\n", &a1, &a2)

	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s1[0] = 100
	fmt.Println(s1, s2)
	fmt.Printf("%p %p\n", &s1, &s2)
	fmt.Printf("%p %p\n", s1, s2)
}
