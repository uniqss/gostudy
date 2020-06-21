package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	fmt.Println("=============================")
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())
}
