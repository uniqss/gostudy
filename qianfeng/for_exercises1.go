package main

import "fmt"

func main() {
	for i := 58; i >= 23; i-- {
		fmt.Println(i)
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	printCount := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t")
			printCount++
			if printCount%5 == 0 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Println()
	fmt.Println(printCount)
}
