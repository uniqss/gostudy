package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	fun1(arr1)
	fmt.Println(arr1)

	s1 := []int{1, 2, 3, 4}
	fmt.Println("slice:", s1)
	fun2(s1)
	fmt.Println("slice:", s1)
}

func fun1(arr [4]int){
	fmt.Println("函数中", arr)
	arr[0] = 100
	fmt.Println("函数中", arr)
}

func fun2(arr []int){
	fmt.Println("函数中 slice:", arr)
	arr[0] = 100
	fmt.Println("函数中 slice:", arr)
}