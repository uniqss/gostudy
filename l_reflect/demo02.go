package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	value := reflect.ValueOf(num)

	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)

	pointer := reflect.ValueOf(&num)
	// panic
	//convertPointer := pointer.Interface().(float64)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)
}
