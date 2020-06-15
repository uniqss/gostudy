package main

import "fmt"

func main() {
	// slice 变长，自动扩容，本质是对现有数组的引用，是引用类型

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)
	fmt.Printf("%T %T\n", arr, s2)

	s3 := make([]int, 3, 8)
	fmt.Println(s3)
	fmt.Printf("容量：%d, 长度：%d\n", cap(s3), len(s3))
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	//fmt.Println(s3[3]) // runtime error: index out of range [3] with length 3

	s4 := make([]int, 0, 5)
	fmt.Println(s4)
	s4 = append(s4, 1, 2)
	fmt.Println(s4)
	s4 = append(s4, 3, 4, 5, 6, 7)
	fmt.Println(s4)

	s4 = append(s4, s3...)
	fmt.Println(s4)

	// 遍历
	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])
	}

	for i, v := range s4 {
		fmt.Printf("%d->%d\n", i, v)
	}
}
