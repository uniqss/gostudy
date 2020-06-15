package main

import "fmt"

func main() {
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)

	for i := 0; i < len(a2); i++ {
		fmt.Printf("&a2[%d]\t%p\n", i, &a2[i])
	}

	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(a2[i]); j++ {
			fmt.Printf("&a2[%d][%d]\t%p\n", i, j, &a2[i][j])
		}
	}

	for _, arr := range a2 {
		for _, val := range arr {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}

	//fmt.Printf("&a2[0][0] %p\n", &a2[0][0])
	//fmt.Printf("&a2[0][1] %p\n", &a2[0][1])
	//fmt.Printf("&a2[0][2] %p\n", &a2[0][2])
	//fmt.Printf("&a2[0][3] %p\n", &a2[0][3])
}
