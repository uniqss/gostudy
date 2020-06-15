package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// math rand
	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	rand.Seed(1)
	num2 := rand.Intn(10)
	fmt.Println("num2:", num2)

	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n", t1)

	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1)
	timeStamp2 := t1.UnixNano()
	fmt.Println(timeStamp2)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println("===>", rand.Intn(100))
	}

	num3 := rand.Intn(46) + 3
	fmt.Println(num3)

	minNum := math.MaxInt32
	maxNum := math.MinInt32
	for i := 1; i < 10000000;i++ {
		r := randRange(50, 100)
		minNum = min(minNum, r)
		maxNum = max(maxNum, r)
	}

	fmt.Println(minNum, maxNum)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// rand [start, stop)
func randRange(start int, stop int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(stop-start) + start
}
