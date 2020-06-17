package main

import "fmt"

func main() {
	a := 10
	fmt.Println("before fun1, a:", a)
	fun1(a)
	fmt.Println("after fun1, a:", a)
	fun2(&a)
	fmt.Println("after fun2, a:", a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("before fun3, arr1:", arr1)
	fun3(arr1)
	fmt.Println("after fun3, arr1:", arr1)
	fun4(&arr1)
	fmt.Println("after fun4, arr1:", arr1)

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("before fun5, s1:", s1)
	fun5(s1)
	fmt.Println("after fun5, s1:", s1)
	fun6(s1)
	fmt.Println("after fun6, s1:", s1)
	fun7(&s1)
	fmt.Println("after fun7, s1:", s1)
}

func fun1(num int) {
	fmt.Println("fun1 num:", num)
	num = 100
	fmt.Println("fun1 ... num:", num)
}
func fun2(p1 *int) {
	fmt.Println("fun2, p1:", *p1)
	*p1 = 200
	fmt.Println("fun2 ... p1:", *p1)
}

func fun3(arr [4]int) {
	fmt.Println("fun3 arr:", arr)
	arr[0] = 1234
	fmt.Println("fun3 ... arr:", arr)
}

func fun4(arr *[4]int) {
	fmt.Println("fun4 arr:", arr)
	arr[0] = 5555
	fmt.Println("fun4 ... arr:", arr)
}

func fun5(s []int) {
	fmt.Println("fun5 s:", s)
	s[0] = 55
	fmt.Println("fun5 ... s:", s)
}

func fun6(s []int) {
	fmt.Println("fun6 s:", s)
	s = append(s, 6, 6, 6, 6, 6, 6)
	fmt.Println("fun6 ... s:", s)
}

func fun7(s *[]int) {
	fmt.Println("fun7 s:", s)
	*s = append(*s, 7, 7, 7, 7, 7, 7, 7)
	fmt.Println("fun7 ... s:", s)
}
