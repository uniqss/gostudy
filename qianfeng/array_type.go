package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [4]float64{1.1, 2.2, 3.3, 4.4}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [2]string{"hello", "world"}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", arr3)
	fmt.Printf("%T\n", arr4)

	num2 := num
	fmt.Println(num, num2)
	num2 = 20
	fmt.Println(num, num2)

	arr5 := arr1
	fmt.Printf("arr1:%p arr5:%p\n", &arr1, &arr5)
	fmt.Println(arr1, arr5)
	arr5[0] = 100
	fmt.Printf("arr1:%p arr5:%p\n", &arr1, &arr5)
	fmt.Println(arr1, arr5)

	fmt.Println(arr5 == arr1)
	arr5[0] = 1
	fmt.Println(arr5 == arr1)

	//fmt.Println(arr1 == arr2) // invalid operation: arr1 == arr2 (mismatched types [4]int and [4]float64)
}
