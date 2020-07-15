package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(arr1[4])

	fmt.Println("======================")
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i*2 + 1
	}
	fmt.Println(arr1)
	fmt.Println("======================")

	for index, value := range arr1 {
		fmt.Printf("下标是:%d数值是%d\n", index, value)
	}

	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println(sum)
}
