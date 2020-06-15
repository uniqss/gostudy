package main

import "fmt"

func main() {
	getSum(1, 2, 3, 4, 5)
	getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func getSum(nums ... int){
	fmt.Printf("%T\n", nums)
	sum := 0
	for i := 0;i < len(nums);i++{
		sum += nums[i]
	}
	fmt.Println(sum)
}
