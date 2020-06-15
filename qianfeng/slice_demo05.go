package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for i:=0;i <len(s1);i++ {
		s2 = append(s2, s1[i])
	}

	fmt.Println(s1, s2)

	s1[0] = 100
	fmt.Println(s1, s2)

	// go 有copy函数，提供给slice做深拷贝
	s3 := []int{7, 8, 9}
	fmt.Println(s2, s3)

	//copy(s2, s3)
	//copy(s3, s2)
	copy(s3[1:], s2[2:])
	fmt.Println(s2, s3)
}
