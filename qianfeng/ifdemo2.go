package main

import "fmt"

func main() {
	if num := -5; num > 0 {
		fmt.Println("正数")
	} else if num < 0 {
		fmt.Println("负数")
	}
}
