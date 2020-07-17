package main

import "fmt"

func main() {
	// for 初始化;判断;递增{}
	// for 判断 {}
	// for {}
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=============== i:", i)

	for {
		fmt.Println("####### i:", i)
	}
}
