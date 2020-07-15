package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 6, 7, 8)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 9, 10)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 11, 12, 13, 14, 15)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)
}
