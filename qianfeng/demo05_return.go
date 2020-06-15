package main

import "fmt"

func main() {
	s1 := getSum1()
	fmt.Println(s1)
	s2 := getSum2()
	fmt.Println(s2)

	res1, res2 := rectangle(5, 3)
	fmt.Println("周长：", res1, " 面积:", res2)

	res3, res4 := rectangle2(5, 3)
	fmt.Println("周长：", res3, " 面积:", res4)

	res5, _ := rectangle2(5, 3)
	fmt.Println(res5)
}

func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func rectangle2(len, wid float64) (perimeter float64, area float64) {
	perimeter = (len + wid) * 2
	area = len * wid
	return
}

func getSum1() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

func getSum2() (sum int) {
	//sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return
}
