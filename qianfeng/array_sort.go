package main

import "fmt"

func main() {
	arr := [5]int{15, 23, 8, 10, 7}

	//冒泡1 思路：每轮把最小的锁定到最前端
	arrlen := len(arr)
	for i := 0; i < arrlen; i++ {
		for j := i + 1; j < arrlen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(arr)

	// 冒泡2 思路：每轮把最大的锁定到最后端
	arr = [5]int{15, 23, 8, 10, 7}
	arrlen = len(arr)
	for i := 1; i < arrlen; i++ {
		for j := 0; j < arrlen-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
