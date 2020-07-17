package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 2; i <= 100; i++ {
		//fmt.Println(i)
		flag := true
		sqrt := int(math.Sqrt(float64(i)))
		for j := 2; j < sqrt; j++ {
			if i%j == 0 {
				//fmt.Printf("%d %% %d == 0\n", i, j)
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
