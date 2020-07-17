package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("已有数据直接创建切片")
	s1 := a[:5]  // 1-5
	s2 := a[3:8] // 4-8
	s3 := a[5:] // 6-10
	s4 := a[:] // 6-10
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)

	fmt.Printf("s1 len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s2 len:%d cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s3 len:%d cap:%d\n", len(s3), cap(s3))
	fmt.Printf("s4 len:%d cap:%d\n", len(s4), cap(s4))
	a[4] = 100
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	s2[2] = 200
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	// 当空间不够时，会重新分配切片，原数组不变
	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	fmt.Printf("s1 len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s2 len:%d cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s3 len:%d cap:%d\n", len(s3), cap(s3))
	fmt.Printf("s4 len:%d cap:%d\n", len(s4), cap(s4))
}
