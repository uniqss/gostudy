package main

import "fmt"

func main() {
	s := getSumRecursive(100)
	fmt.Println(s)

	f := getFeibonacci(10)
	fmt.Println(f)
}

func getSumRecursive(num int) int {
	if num == 1 {
		return 1
	}
	return num + getSumRecursive(num - 1)
}

func getFeibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return getFeibonacci(n - 1) + getFeibonacci(n - 2)
}