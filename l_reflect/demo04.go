package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	fmt.Println("num的数值是：", num)

	ptr := reflect.ValueOf(&num) // 注意参数必须是指针，取地址
	newValue := ptr.Elem()

	fmt.Println("类型：", newValue.Type())
	fmt.Println("是否可以修改类型：", newValue.CanSet())

	newValue.SetFloat(3.14)
	fmt.Println(num)
}
