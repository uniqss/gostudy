package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("------------------------------")

	out:for i := 1;i <= 5;i++ {
		for j := 1;j <= 5;j++ {
			if j == 2 {
				break out
			}
			fmt.Printf("i:%d, j:%d\n", i, j)
		}
	}

	fmt.Println("main over")
}
