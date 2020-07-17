package main

import "fmt"

func main() {
	s1 := getSum(1, 10)
	fmt.Printf("%d\n", s1)
}

func getSum(begin, end int) int {
	sum := 0
	for i:=begin;i<=end;i++ {
		sum += i
	}
	return sum
}
