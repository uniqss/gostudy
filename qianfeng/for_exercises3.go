package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 100; i < 1000; i++ {
		x := i / 100
		y := i / 10 % 10
		z := i % 10

		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}

	fmt.Println("==========")
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				abc := a*100 + b*10 + c
				if math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3) == float64(abc) {
					fmt.Println(abc)
				}
			}
		}
	}
}
